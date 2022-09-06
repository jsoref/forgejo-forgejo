// SPDX-License-Identifier: MIT

package driver

import (
	"context"
	"fmt"
	"time"

	"code.gitea.io/gitea/models/db"
	issues_model "code.gitea.io/gitea/models/issues"
	repo_model "code.gitea.io/gitea/models/repo"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/git"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/timeutil"
	issue_service "code.gitea.io/gitea/services/issue"

	f3_forgejo "lab.forgefriends.org/friendlyforgeformat/gof3/forges/forgejo"
	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
	"lab.forgefriends.org/friendlyforgeformat/gof3/util"
)

type PullRequest struct {
	issues_model.PullRequest
	FetchFunc func(repository string) string
}

func PullRequestConverter(f *issues_model.PullRequest) *PullRequest {
	return &PullRequest{
		PullRequest: *f,
	}
}

func (o PullRequest) GetID() int64 {
	return o.Index
}

func (o PullRequest) GetIDString() string {
	return fmt.Sprintf("%d", o.GetID())
}

func (o *PullRequest) SetID(id int64) {
	o.Index = id
}

func (o *PullRequest) SetIDString(id string) {
	o.SetID(util.ParseInt(id))
}

func (o *PullRequest) IsNil() bool {
	return o.Index == 0
}

func (o *PullRequest) Equals(other *PullRequest) bool {
	return o.Issue.Title == other.Issue.Title
}

func (o PullRequest) IsForkPullRequest() bool {
	return o.HeadRepoID != o.BaseRepoID
}

func (o *PullRequest) ToFormatInterface() format.Interface {
	return o.ToFormat()
}

func (o *PullRequest) ToFormat() *format.PullRequest {
	var milestone string
	if o.Issue.Milestone != nil {
		milestone = o.Issue.Milestone.Name
	}

	labels := make([]string, 0, len(o.Issue.Labels))
	for _, label := range o.Issue.Labels {
		labels = append(labels, label.Name)
	}

	var mergedTime *time.Time
	if o.HasMerged {
		mergedTime = o.MergedUnix.AsTimePtr()
	}

	getSHA := func(repo *repo_model.Repository, branch string) string {
		r, err := git.OpenRepository(context.Background(), repo.RepoPath())
		if err != nil {
			panic(err)
		}
		defer r.Close()

		b, err := r.GetBranch(branch)
		if err != nil {
			panic(err)
		}

		c, err := b.GetCommit()
		if err != nil {
			panic(err)
		}
		return c.ID.String()
	}

	head := format.PullRequestBranch{
		CloneURL:  o.HeadRepo.CloneLink().HTTPS,
		Ref:       o.HeadBranch,
		SHA:       getSHA(o.HeadRepo, o.HeadBranch),
		RepoName:  o.HeadRepo.Name,
		OwnerName: o.HeadRepo.OwnerName,
	}

	base := format.PullRequestBranch{
		CloneURL:  o.BaseRepo.CloneLink().HTTPS,
		Ref:       o.BaseBranch,
		SHA:       getSHA(o.BaseRepo, o.BaseBranch),
		RepoName:  o.BaseRepo.Name,
		OwnerName: o.BaseRepo.OwnerName,
	}

	return &format.PullRequest{
		Common:         format.NewCommon(o.Index),
		PosterID:       format.NewUserReference(o.Issue.Poster.ID),
		Title:          o.Issue.Title,
		Content:        o.Issue.Content,
		Milestone:      milestone,
		State:          string(o.Issue.State()),
		IsLocked:       o.Issue.IsLocked,
		Created:        o.Issue.CreatedUnix.AsTime(),
		Updated:        o.Issue.UpdatedUnix.AsTime(),
		Closed:         o.Issue.ClosedUnix.AsTimePtr(),
		Labels:         labels,
		PatchURL:       o.Issue.PatchURL(),
		Merged:         o.HasMerged,
		MergedTime:     mergedTime,
		MergeCommitSHA: o.MergedCommitID,
		Head:           head,
		Base:           base,
	}
}

func (o *PullRequest) FromFormat(pullRequest *format.PullRequest) {
	labels := make([]*issues_model.Label, 0, len(pullRequest.Labels))
	for _, label := range pullRequest.Labels {
		labels = append(labels, &issues_model.Label{Name: label})
	}

	if pullRequest.Created.IsZero() {
		if pullRequest.Closed != nil {
			pullRequest.Created = *pullRequest.Closed
		} else if pullRequest.MergedTime != nil {
			pullRequest.Created = *pullRequest.MergedTime
		} else {
			pullRequest.Created = time.Now()
		}
	}
	if pullRequest.Updated.IsZero() {
		pullRequest.Updated = pullRequest.Created
	}

	ctx := context.Background()
	base, err := repo_model.GetRepositoryByOwnerAndName(ctx, pullRequest.Base.OwnerName, pullRequest.Base.RepoName)
	if err != nil {
		panic(err)
	}
	var head *repo_model.Repository
	if pullRequest.Head.RepoName == "" {
		head = base
	} else {
		head, err = repo_model.GetRepositoryByOwnerAndName(ctx, pullRequest.Head.OwnerName, pullRequest.Head.RepoName)
		if err != nil {
			panic(err)
		}
	}

	issue := issues_model.Issue{
		RepoID:   base.ID,
		Repo:     base,
		Title:    pullRequest.Title,
		Index:    pullRequest.GetID(),
		PosterID: pullRequest.PosterID.GetID(),
		Poster: &user_model.User{
			ID: pullRequest.PosterID.GetID(),
		},
		Content:     pullRequest.Content,
		IsPull:      true,
		IsClosed:    pullRequest.State == "closed",
		IsLocked:    pullRequest.IsLocked,
		Labels:      labels,
		CreatedUnix: timeutil.TimeStamp(pullRequest.Created.Unix()),
		UpdatedUnix: timeutil.TimeStamp(pullRequest.Updated.Unix()),
	}

	pr := issues_model.PullRequest{
		HeadRepoID: head.ID,
		HeadRepo: &repo_model.Repository{
			ID:        head.ID,
			Name:      pullRequest.Head.RepoName,
			OwnerName: pullRequest.Head.OwnerName,
		},
		HeadBranch: pullRequest.Head.Ref,
		BaseRepoID: base.ID,
		BaseRepo: &repo_model.Repository{
			ID:        base.ID,
			Name:      pullRequest.Base.RepoName,
			OwnerName: pullRequest.Base.OwnerName,
		},
		BaseBranch: pullRequest.Base.Ref,
		MergeBase:  pullRequest.Base.SHA,
		Index:      pullRequest.GetID(),
		HasMerged:  pullRequest.Merged,

		Issue: &issue,
	}

	if pr.Issue.IsClosed && pullRequest.Closed != nil {
		pr.Issue.ClosedUnix = timeutil.TimeStamp(pullRequest.Closed.Unix())
	}
	if pr.HasMerged && pullRequest.MergedTime != nil {
		pr.MergedUnix = timeutil.TimeStamp(pullRequest.MergedTime.Unix())
		pr.MergedCommitID = pullRequest.MergeCommitSHA
	}

	*o = PullRequest{
		PullRequest: pr,
		FetchFunc:   pullRequest.FetchFunc,
	}
}

type PullRequestProvider struct {
	BaseProviderWithProjectProvider
	prHeadCache f3_forgejo.PrHeadCache
}

func (o *PullRequestProvider) ToFormat(ctx context.Context, pullRequest *PullRequest) *format.PullRequest {
	return pullRequest.ToFormat()
}

func (o *PullRequestProvider) FromFormat(ctx context.Context, pr *format.PullRequest) *PullRequest {
	var pullRequest PullRequest
	pullRequest.FromFormat(pr)
	return &pullRequest
}

func (o *PullRequestProvider) Init() *PullRequestProvider {
	o.prHeadCache = make(f3_forgejo.PrHeadCache)
	return o
}

func (o *PullRequestProvider) cleanupRemotes(ctx context.Context, repository string) {
	for remote := range o.prHeadCache {
		util.Command(ctx, "git", "-C", repository, "remote", "rm", remote)
	}
	o.prHeadCache = make(f3_forgejo.PrHeadCache)
}

func (o *PullRequestProvider) GetObjects(ctx context.Context, user *User, project *Project, page int) []*PullRequest {
	pullRequests, _, err := issues_model.PullRequests(ctx, project.GetID(), &issues_model.PullRequestsOptions{
		ListOptions: db.ListOptions{Page: page, PageSize: o.g.perPage},
		State:       string(api.StateAll),
	})
	if err != nil {
		panic(fmt.Errorf("error while listing pullRequests: %v", err))
	}

	return util.ConvertMap[*issues_model.PullRequest, *PullRequest](pullRequests, PullRequestConverter)
}

func (o *PullRequestProvider) ProcessObject(ctx context.Context, user *User, project *Project, pr *PullRequest) {
	if err := pr.LoadIssue(ctx); err != nil {
		panic(err)
	}
	if err := pr.Issue.LoadRepo(ctx); err != nil {
		panic(err)
	}
	if err := pr.Issue.LoadAttributes(ctx); err != nil {
		panic(err)
	}
	if err := pr.LoadAttributes(ctx); err != nil {
		panic(err)
	}
	if err := pr.LoadBaseRepo(ctx); err != nil {
		panic(err)
	}
	if err := pr.LoadHeadRepo(ctx); err != nil {
		panic(err)
	}

	pr.FetchFunc = func(repository string) string {
		head, messages := f3_forgejo.UpdateGitForPullRequest(ctx, &o.prHeadCache, pr.ToFormat(), repository)
		for _, message := range messages {
			o.g.GetLogger().Warn(message)
		}
		o.cleanupRemotes(ctx, repository)
		return head
	}
}

func (o *PullRequestProvider) Get(ctx context.Context, user *User, project *Project, pullRequest *PullRequest) *PullRequest {
	id := pullRequest.GetID()
	pr, err := issues_model.GetPullRequestByIndex(ctx, project.GetID(), id)
	if issues_model.IsErrPullRequestNotExist(err) {
		return &PullRequest{}
	}
	if err != nil {
		panic(err)
	}
	p := PullRequestConverter(pr)
	o.ProcessObject(ctx, user, project, p)
	return p
}

func (o *PullRequestProvider) Put(ctx context.Context, user *User, project *Project, pullRequest, existing *PullRequest) *PullRequest {
	i := pullRequest.PullRequest.Issue
	i.RepoID = project.GetID()
	labels := make([]int64, 0, len(i.Labels))
	for _, label := range i.Labels {
		labels = append(labels, label.ID)
	}

	if existing == nil || existing.IsNil() {
		if err := issues_model.NewPullRequest(ctx, &project.Repository, i, labels, []string{}, &pullRequest.PullRequest); err != nil {
			panic(err)
		}
	} else {
		var u issues_model.Issue
		u.Index = i.Index
		u.RepoID = project.GetID()
		cols := make([]string, 0, 10)

		if i.Title != existing.Issue.Title {
			u.Title = i.Title
			cols = append(cols, "name")
		}

		if len(cols) > 0 {
			if _, err := db.GetEngine(ctx).ID(existing.Issue.ID).Cols(cols...).Update(u); err != nil {
				panic(err)
			}
		}
	}

	if pullRequest.FetchFunc != nil {
		repoPath := repo_model.RepoPath(user.Name, project.Name)
		fromHead := pullRequest.FetchFunc(repoPath)
		gitRepo, err := git.OpenRepository(ctx, repoPath)
		if err != nil {
			panic(err)
		}
		defer gitRepo.Close()

		toHead := fmt.Sprintf("%s%d/head", git.PullPrefix, pullRequest.GetID())
		if err := git.NewCommand(ctx, "update-ref").AddDynamicArguments(toHead, fromHead).Run(&git.RunOpts{Dir: repoPath}); err != nil {
			panic(err)
		}
	}

	return o.Get(ctx, user, project, pullRequest)
}

func (o *PullRequestProvider) Delete(ctx context.Context, user *User, project *Project, pullRequest *PullRequest) *PullRequest {
	p := o.Get(ctx, user, project, pullRequest)
	if !p.IsNil() {
		repoPath := repo_model.RepoPath(user.Name, project.Name)
		gitRepo, err := git.OpenRepository(ctx, repoPath)
		if err != nil {
			panic(err)
		}
		defer gitRepo.Close()
		if err := issue_service.DeleteIssue(ctx, o.g.GetDoer(), gitRepo, p.PullRequest.Issue); err != nil {
			panic(err)
		}
	}
	return p
}
