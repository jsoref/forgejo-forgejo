// SPDX-License-Identifier: MIT

package driver

import (
	"context"

	repo_model "code.gitea.io/gitea/models/repo"
	base "code.gitea.io/gitea/modules/migration"
	repo_module "code.gitea.io/gitea/modules/repository"
	"code.gitea.io/gitea/services/migrations"

	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
	"lab.forgefriends.org/friendlyforgeformat/gof3/util"
)

type Repository struct {
	format.Repository
}

func (o *Repository) Equals(other *Repository) bool {
	return false // it is costly to figure that out, mirroring is as fast
}

func (o *Repository) ToFormatInterface() format.Interface {
	return o.ToFormat()
}

func (o *Repository) ToFormat() *format.Repository {
	return &o.Repository
}

func (o *Repository) FromFormat(repository *format.Repository) {
	o.Repository = *repository
}

type RepositoryProvider struct {
	BaseProvider
}

func (o *RepositoryProvider) ToFormat(ctx context.Context, repository *Repository) *format.Repository {
	return repository.ToFormat()
}

func (o *RepositoryProvider) FromFormat(ctx context.Context, p *format.Repository) *Repository {
	var repository Repository
	repository.FromFormat(p)
	return &repository
}

func (o *RepositoryProvider) GetObjects(ctx context.Context, user *User, project *Project, page int) []*Repository {
	if page > 1 {
		return make([]*Repository, 0)
	}
	repositories := make([]*Repository, 0, len(format.RepositoryNames))
	for _, name := range format.RepositoryNames {
		repositories = append(repositories, o.Get(ctx, user, project, &Repository{
			Repository: format.Repository{
				Name: name,
			},
		}))
	}
	return repositories
}

func (o *RepositoryProvider) ProcessObject(ctx context.Context, user *User, project *Project, repository *Repository) {
}

func (o *RepositoryProvider) Get(ctx context.Context, user *User, project *Project, exemplar *Repository) *Repository {
	repoPath := repo_model.RepoPath(user.Name, project.Name) + exemplar.Name
	o.g.GetLogger().Debug(repoPath)
	return &Repository{
		Repository: format.Repository{
			Name: exemplar.Name,
			FetchFunc: func(destination string) {
				o.g.GetLogger().Debug("RepositoryProvider:Get: git clone %s %s", repoPath, destination)
				util.Command(ctx, "git", "clone", "--mirror", repoPath, destination)
			},
		},
	}
}

func (o *RepositoryProvider) Put(ctx context.Context, user *User, project *Project, repository, existing *Repository) *Repository {
	if repository.FetchFunc != nil {
		directory, delete := format.RepositoryDefaultDirectory()
		defer delete()
		repository.FetchFunc(directory)

		_, err := repo_module.MigrateRepositoryGitData(ctx, &user.User, &project.Repository, base.MigrateOptions{
			RepoName:       project.Name,
			Mirror:         false,
			MirrorInterval: "",
			LFS:            false,
			LFSEndpoint:    "",
			CloneAddr:      directory,
			Wiki:           o.g.GetOptions().GetFeatures().Wiki,
			Releases:       o.g.GetOptions().GetFeatures().Releases,
		}, migrations.NewMigrationHTTPTransport())
		if err != nil {
			panic(err)
		}
	}
	return o.Get(ctx, user, project, repository)
}

func (o *RepositoryProvider) Delete(ctx context.Context, user *User, project *Project, repository *Repository) *Repository {
	panic("It is not possible to delete a repository")
}
