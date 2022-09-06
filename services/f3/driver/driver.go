// SPDX-License-Identifier: MIT

package driver

import (
	"context"
	"fmt"

	auth_model "code.gitea.io/gitea/models/auth"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/migrations"

	"github.com/urfave/cli/v2"
	config_factory "lab.forgefriends.org/friendlyforgeformat/gof3/config/factory"
	f3_types "lab.forgefriends.org/friendlyforgeformat/gof3/config/types"
	"lab.forgefriends.org/friendlyforgeformat/gof3/forges/common"
	"lab.forgefriends.org/friendlyforgeformat/gof3/forges/driver"
	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
)

var Name = "InternalForgejo"

func init() {
	config_factory.RegisterFactory(Name, f3_types.OptionsFactory{
		Name:  Name,
		New:   func() f3_types.OptionsInterface { return &Options{} },
		Flags: GetFlags,
	}, func() common.DriverInterface { return &Forgejo{} })
}

type Options struct {
	f3_types.Options

	AuthenticationSource int64
	Doer                 *user_model.User
}

func getAuthenticationSource(ctx context.Context, authenticationSource string) (*auth_model.Source, error) {
	source, err := auth_model.GetSourceByName(ctx, authenticationSource)
	if err != nil {
		if auth_model.IsErrSourceNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return source, nil
}

func (o *Options) FromFlags(ctx context.Context, c *cli.Context, prefix string) f3_types.OptionsInterface {
	o.Options.FromFlags(ctx, c, prefix)
	sourceName := c.String("authentication-source")
	if sourceName != "" {
		source, err := getAuthenticationSource(ctx, sourceName)
		if err != nil {
			panic(fmt.Errorf("error retrieving the authentication-source %s %v", sourceName, err))
		}
		if source != nil {
			o.AuthenticationSource = source.ID
		}
	}

	doer, err := user_model.GetAdminUser(ctx)
	if err != nil {
		panic(fmt.Errorf("GetAdminUser %v", err))
	}
	o.Doer = doer

	return o
}

func GetFlags(prefix, category string) []cli.Flag {
	flags := make([]cli.Flag, 0, 10)

	flags = append(flags, &cli.StringFlag{
		Name:  "authentication-source",
		Value: "",
		Usage: "The name of the authentication source matching the forge of origin",
	})

	return flags
}

type Forgejo struct {
	perPage int
	options *Options
}

func (o *Forgejo) GetName() string {
	return Name
}

func (o *Forgejo) GetPerPage() int {
	return o.perPage
}

func (o *Forgejo) GetOptions() f3_types.OptionsInterface {
	return o.options
}

func (o *Forgejo) SetOptions(options f3_types.OptionsInterface) {
	var ok bool
	o.options, ok = options.(*Options)
	if !ok {
		panic(fmt.Errorf("unexpected type %T", options))
	}
}

func (o *Forgejo) GetLogger() *f3_types.Logger {
	return o.GetOptions().GetLogger()
}

func (o *Forgejo) Init(options f3_types.OptionsInterface) {
	o.SetOptions(options)
	o.perPage = setting.ItemsPerPage
}

func (o *Forgejo) GetDirectory() string {
	return o.options.GetDirectory()
}

func (o *Forgejo) GetDoer() *user_model.User {
	return o.options.Doer
}

func (o *Forgejo) GetAuthenticationSource() int64 {
	return o.options.AuthenticationSource
}

func (o *Forgejo) GetNewMigrationHTTPClient() f3_types.NewMigrationHTTPClientFun {
	return migrations.NewMigrationHTTPClient
}

func (o *Forgejo) SupportGetRepoComments() bool {
	return false
}

func (o *Forgejo) GetProvider(name string, parent common.ProviderInterface) common.ProviderInterface {
	var parentImpl any
	if parent != nil {
		parentImpl = parent.GetImplementation()
	}
	switch name {
	case driver.ProviderUser:
		return driver.NewProvider[UserProvider, *UserProvider, User, *User, format.User, *format.User](driver.ProviderUser, NewProvider[UserProvider](o))
	case driver.ProviderProject:
		return driver.NewProviderWithParentOne[ProjectProvider, *ProjectProvider, Project, *Project, format.Project, *format.Project, User, *User](driver.ProviderProject, NewProvider[ProjectProvider, *ProjectProvider](o))
	case driver.ProviderMilestone:
		return driver.NewProviderWithParentOneTwo[MilestoneProvider, *MilestoneProvider, Milestone, *Milestone, format.Milestone, *format.Milestone, User, *User, Project, *Project](driver.ProviderMilestone, NewProviderWithProjectProvider[MilestoneProvider](o, parentImpl.(*ProjectProvider)))
	case driver.ProviderIssue:
		return driver.NewProviderWithParentOneTwo[IssueProvider, *IssueProvider, Issue, *Issue, format.Issue, *format.Issue, User, *User, Project, *Project](driver.ProviderIssue, NewProviderWithProjectProvider[IssueProvider](o, parentImpl.(*ProjectProvider)))
	case driver.ProviderPullRequest:
		return driver.NewProviderWithParentOneTwo[PullRequestProvider, *PullRequestProvider, PullRequest, *PullRequest, format.PullRequest, *format.PullRequest, User, *User, Project, *Project](driver.ProviderPullRequest, NewProviderWithProjectProvider[PullRequestProvider](o, parentImpl.(*ProjectProvider)))
	case driver.ProviderReview:
		return driver.NewProviderWithParentOneTwoThree[ReviewProvider, *ReviewProvider, Review, *Review, format.Review, *format.Review, User, *User, Project, *Project, PullRequest, *PullRequest](driver.ProviderReview, NewProvider[ReviewProvider](o))
	case driver.ProviderRepository:
		return driver.NewProviderWithParentOneTwo[RepositoryProvider, *RepositoryProvider, Repository, *Repository, format.Repository, *format.Repository, User, *User, Project, *Project](driver.ProviderRepository, NewProvider[RepositoryProvider](o))
	case driver.ProviderTopic:
		return driver.NewProviderWithParentOneTwo[TopicProvider, *TopicProvider, Topic, *Topic, format.Topic, *format.Topic, User, *User, Project, *Project](driver.ProviderTopic, NewProvider[TopicProvider](o))
	case driver.ProviderLabel:
		return driver.NewProviderWithParentOneTwo[LabelProvider, *LabelProvider, Label, *Label, format.Label, *format.Label, User, *User, Project, *Project](driver.ProviderLabel, NewProviderWithProjectProvider[LabelProvider](o, parentImpl.(*ProjectProvider)))
	case driver.ProviderRelease:
		return driver.NewProviderWithParentOneTwo[ReleaseProvider, *ReleaseProvider, Release, *Release, format.Release, *format.Release, User, *User, Project, *Project](driver.ProviderRelease, NewProvider[ReleaseProvider](o))
	case driver.ProviderAsset:
		return driver.NewProviderWithParentOneTwoThree[AssetProvider, *AssetProvider, Asset, *Asset, format.ReleaseAsset, *format.ReleaseAsset, User, *User, Project, *Project, Release, *Release](driver.ProviderAsset, NewProvider[AssetProvider](o))
	case driver.ProviderComment:
		return driver.NewProviderWithParentOneTwoThreeInterface[CommentProvider, *CommentProvider, Comment, *Comment, format.Comment, *format.Comment, User, *User, Project, *Project](driver.ProviderComment, NewProvider[CommentProvider](o))
	case driver.ProviderReaction:
		return driver.NewProviderWithParentOneTwoRest[ReactionProvider, *ReactionProvider, Reaction, *Reaction, format.Reaction, *format.Reaction, User, *User, Project, *Project](driver.ProviderReaction, NewProvider[ReactionProvider](o))
	default:
		panic(fmt.Sprintf("unknown provider name %s", name))
	}
}

func (o Forgejo) Finish() {
}
