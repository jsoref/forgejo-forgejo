// SPDX-License-Identifier: MIT

package driver

import (
	"context"
	"fmt"

	"code.gitea.io/gitea/models/db"
	issues_model "code.gitea.io/gitea/models/issues"
	repo_model "code.gitea.io/gitea/models/repo"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/timeutil"
	issue_service "code.gitea.io/gitea/services/issue"

	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
	"lab.forgefriends.org/friendlyforgeformat/gof3/util"
)

type Issue struct {
	issues_model.Issue
}

func IssueConverter(f *issues_model.Issue) *Issue {
	return &Issue{
		Issue: *f,
	}
}

func (o Issue) GetID() int64 {
	return o.Index
}

func (o Issue) GetIDString() string {
	return fmt.Sprintf("%d", o.GetID())
}

func (o *Issue) SetID(id int64) {
	o.Index = id
}

func (o *Issue) SetIDString(id string) {
	o.SetID(util.ParseInt(id))
}

func (o *Issue) IsNil() bool {
	return o.Index == 0
}

func (o *Issue) Equals(other *Issue) bool {
	return (o.Title == other.Title)
}

func (o *Issue) ToFormatInterface() format.Interface {
	return o.ToFormat()
}

func (o *Issue) ToFormat() *format.Issue {
	var milestone string
	if o.Milestone != nil {
		milestone = o.Milestone.Name
	}

	labels := make([]string, 0, len(o.Labels))
	for _, label := range o.Labels {
		labels = append(labels, label.Name)
	}

	var assignees []string
	for i := range o.Assignees {
		assignees = append(assignees, o.Assignees[i].Name)
	}

	return &format.Issue{
		Common:    format.NewCommon(o.Index),
		Title:     o.Title,
		PosterID:  format.NewUserReference(o.Poster.ID),
		Content:   o.Content,
		Milestone: milestone,
		State:     string(o.State()),
		Created:   o.CreatedUnix.AsTime(),
		Updated:   o.UpdatedUnix.AsTime(),
		Closed:    o.ClosedUnix.AsTimePtr(),
		IsLocked:  o.IsLocked,
		Labels:    labels,
		Assignees: assignees,
	}
}

func (o *Issue) FromFormat(issue *format.Issue) {
	labels := make([]*issues_model.Label, 0, len(issue.Labels))
	for _, label := range issue.Labels {
		labels = append(labels, &issues_model.Label{Name: label})
	}

	assignees := make([]*user_model.User, 0, len(issue.Assignees))
	for _, a := range issue.Assignees {
		assignees = append(assignees, &user_model.User{
			Name: a,
		})
	}

	*o = Issue{
		Issue: issues_model.Issue{
			Title:    issue.Title,
			Index:    issue.GetID(),
			PosterID: issue.PosterID.GetID(),
			Poster: &user_model.User{
				ID: issue.PosterID.GetID(),
			},
			Content: issue.Content,
			Milestone: &issues_model.Milestone{
				Name: issue.Milestone,
			},
			IsClosed:    issue.State == "closed",
			CreatedUnix: timeutil.TimeStamp(issue.Created.Unix()),
			UpdatedUnix: timeutil.TimeStamp(issue.Updated.Unix()),
			IsLocked:    issue.IsLocked,
			Labels:      labels,
			Assignees:   assignees,
		},
	}

	if issue.Closed != nil {
		o.ClosedUnix = timeutil.TimeStamp(issue.Closed.Unix())
	}
}

type IssueProvider struct {
	BaseProviderWithProjectProvider
}

func (o *IssueProvider) ToFormat(ctx context.Context, issue *Issue) *format.Issue {
	return issue.ToFormat()
}

func (o *IssueProvider) FromFormat(ctx context.Context, i *format.Issue) *Issue {
	var issue Issue
	issue.FromFormat(i)
	if i.Milestone != "" {
		issue.Milestone.ID = o.project.milestones.GetID(issue.Milestone.Name)
	}
	for _, label := range issue.Labels {
		label.ID = o.project.labels.GetID(label.Name)
	}
	return &issue
}

func (o *IssueProvider) GetObjects(ctx context.Context, user *User, project *Project, page int) []*Issue {
	issues, err := issues_model.Issues(ctx, &issues_model.IssuesOptions{
		Paginator: &db.ListOptions{Page: page, PageSize: o.g.perPage},
		RepoIDs:   []int64{project.GetID()},
	})
	if err != nil {
		panic(fmt.Errorf("error while listing issues: %v", err))
	}

	return util.ConvertMap[*issues_model.Issue, *Issue](issues, IssueConverter)
}

func (o *IssueProvider) ProcessObject(ctx context.Context, user *User, project *Project, issue *Issue) {
	if err := (&issue.Issue).LoadAttributes(ctx); err != nil {
		panic(true)
	}
}

func (o *IssueProvider) Get(ctx context.Context, user *User, project *Project, exemplar *Issue) *Issue {
	id := exemplar.GetID()
	issue, err := issues_model.GetIssueByIndex(ctx, project.GetID(), id)
	if issues_model.IsErrIssueNotExist(err) {
		return &Issue{}
	}
	if err != nil {
		panic(err)
	}
	i := IssueConverter(issue)
	o.ProcessObject(ctx, user, project, i)
	return i
}

func (o *IssueProvider) Put(ctx context.Context, user *User, project *Project, issue, existing *Issue) *Issue {
	i := issue.Issue
	i.RepoID = project.GetID()
	makeLabels := func(issueID int64) []issues_model.IssueLabel {
		labels := make([]issues_model.IssueLabel, 0, len(i.Labels))
		for _, label := range i.Labels {
			o.g.GetLogger().Trace("%d with label %d", issueID, label.ID)
			labels = append(labels, issues_model.IssueLabel{
				IssueID: issueID,
				LabelID: label.ID,
			})
		}
		return labels
	}

	var result *Issue

	sess := db.GetEngine(ctx)
	if existing == nil || existing.IsNil() {
		idx, err := db.GetNextResourceIndex(ctx, "issue_index", project.Repository.ID)
		if err != nil {
			panic(fmt.Errorf("generate issue index failed: %w", err))
		}
		i.Index = idx

		if _, err = sess.NoAutoTime().Insert(&i); err != nil {
			panic(err)
		}
		labels := makeLabels(i.ID)
		if len(labels) > 0 {
			if _, err := sess.Insert(labels); err != nil {
				panic(err)
			}
		}
		result = IssueConverter(&i)
	} else {
		e := existing.Issue
		if issue.GetID() != existing.GetID() {
			panic(fmt.Sprintf("issue.GetID() %d != existing.GetID() %d", issue.GetID(), existing.GetID()))
		}
		var u issues_model.Issue
		u.Index = existing.GetID()
		u.RepoID = project.GetID()
		cols := make([]string, 0, 10)
		if i.Title != e.Title {
			u.Title = i.Title
			cols = append(cols, "name")
		}

		if len(cols) > 0 {
			if _, err := sess.ID(existing.ID).Cols(cols...).Update(u); err != nil {
				panic(err)
			}
		}

		result = existing
	}

	return o.Get(ctx, user, project, result)
}

func (o *IssueProvider) Delete(ctx context.Context, user *User, project *Project, issue *Issue) *Issue {
	m := o.Get(ctx, user, project, issue)
	if !m.IsNil() {
		repoPath := repo_model.RepoPath(user.Name, project.Name)
		gitRepo, err := git.OpenRepository(ctx, repoPath)
		if err != nil {
			panic(err)
		}
		defer gitRepo.Close()
		if err := issue_service.DeleteIssue(ctx, o.g.GetDoer(), gitRepo, &issue.Issue); err != nil {
			panic(err)
		}
	}
	return m
}
