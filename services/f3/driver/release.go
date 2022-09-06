// SPDX-License-Identifier: MIT

package driver

import (
	"context"
	"fmt"
	"strings"
	"time"

	"code.gitea.io/gitea/models/db"
	repo_model "code.gitea.io/gitea/models/repo"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/timeutil"
	release_service "code.gitea.io/gitea/services/release"

	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
	"lab.forgefriends.org/friendlyforgeformat/gof3/util"
)

type Release struct {
	repo_model.Release
}

func ReleaseConverter(f *repo_model.Release) *Release {
	return &Release{
		Release: *f,
	}
}

func (o Release) GetID() int64 {
	return o.ID
}

func (o Release) GetIDString() string {
	return fmt.Sprintf("%d", o.GetID())
}

func (o *Release) SetID(id int64) {
	o.ID = id
}

func (o *Release) SetIDString(id string) {
	o.SetID(util.ParseInt(id))
}

func (o *Release) IsNil() bool {
	return o.ID == 0
}

func (o *Release) Equals(other *Release) bool {
	return (o.TagName == other.TagName &&
		o.Title == other.Title)
}

func (o *Release) ToFormatInterface() format.Interface {
	return o.ToFormat()
}

func (o *Release) ToFormat() *format.Release {
	return &format.Release{
		Common:          format.NewCommon(o.ID),
		TagName:         o.TagName,
		TargetCommitish: o.Target,
		Name:            o.Title,
		Body:            o.Note,
		Draft:           o.IsDraft,
		Prerelease:      o.IsPrerelease,
		Created:         o.CreatedUnix.AsTime(),
		PublisherID:     format.NewUserReference(o.Publisher.ID),
	}
}

func (o *Release) FromFormat(release *format.Release) {
	if release.Created.IsZero() {
		if !release.Published.IsZero() {
			release.Created = release.Published
		} else {
			release.Created = time.Now()
		}
	}

	*o = Release{
		repo_model.Release{
			ID:          release.GetID(),
			PublisherID: release.PublisherID.GetID(),
			Publisher: &user_model.User{
				ID: release.PublisherID.GetID(),
			},
			TagName:      release.TagName,
			LowerTagName: strings.ToLower(release.TagName),
			Target:       release.TargetCommitish,
			Title:        release.Name,
			Note:         release.Body,
			IsDraft:      release.Draft,
			IsPrerelease: release.Prerelease,
			IsTag:        false,
			CreatedUnix:  timeutil.TimeStamp(release.Created.Unix()),
		},
	}
}

type ReleaseProvider struct {
	BaseProvider
}

func (o *ReleaseProvider) ToFormat(ctx context.Context, release *Release) *format.Release {
	return release.ToFormat()
}

func (o *ReleaseProvider) FromFormat(ctx context.Context, i *format.Release) *Release {
	var release Release
	release.FromFormat(i)
	return &release
}

func (o *ReleaseProvider) GetObjects(ctx context.Context, user *User, project *Project, page int) []*Release {
	releases, err := repo_model.GetReleasesByRepoID(ctx, project.GetID(), repo_model.FindReleasesOptions{
		ListOptions:   db.ListOptions{Page: page, PageSize: o.g.perPage},
		IncludeDrafts: true,
		IncludeTags:   false,
	})
	if err != nil {
		panic(fmt.Errorf("error while listing releases: %v", err))
	}

	return util.ConvertMap[*repo_model.Release, *Release](releases, ReleaseConverter)
}

func (o *ReleaseProvider) ProcessObject(ctx context.Context, user *User, project *Project, release *Release) {
	if err := (&release.Release).LoadAttributes(ctx); err != nil {
		panic(err)
	}
}

func (o *ReleaseProvider) Get(ctx context.Context, user *User, project *Project, exemplar *Release) *Release {
	id := exemplar.GetID()
	release, err := repo_model.GetReleaseByID(ctx, id)
	if repo_model.IsErrReleaseNotExist(err) {
		return &Release{}
	}
	if err != nil {
		panic(err)
	}
	r := ReleaseConverter(release)
	o.ProcessObject(ctx, user, project, r)
	return r
}

func (o *ReleaseProvider) Put(ctx context.Context, user *User, project *Project, release, existing *Release) *Release {
	var result *Release

	if existing == nil || existing.IsNil() {
		r := release.Release
		r.RepoID = project.GetID()

		repoPath := repo_model.RepoPath(user.Name, project.Name)
		gitRepo, err := git.OpenRepository(ctx, repoPath)
		if err != nil {
			panic(err)
		}
		defer gitRepo.Close()

		if err := release_service.CreateRelease(gitRepo, &r, nil, ""); err != nil {
			panic(err)
		}
		result = ReleaseConverter(&r)
	} else {
		var u repo_model.Release
		u.ID = existing.GetID()
		cols := make([]string, 0, 10)

		if release.Title != existing.Title {
			u.Title = release.Title
			cols = append(cols, "title")
		}
		if len(cols) > 0 {
			if _, err := db.GetEngine(ctx).ID(existing.ID).Cols(cols...).Update(u); err != nil {
				panic(err)
			}
		}
		result = existing
	}
	return o.Get(ctx, user, project, result)
}

func (o *ReleaseProvider) Delete(ctx context.Context, user *User, project *Project, release *Release) *Release {
	m := o.Get(ctx, user, project, release)
	if !m.IsNil() {
		if err := release_service.DeleteReleaseByID(ctx, release.GetID(), o.g.GetDoer(), false); err != nil {
			panic(err)
		}
	}
	return m
}
