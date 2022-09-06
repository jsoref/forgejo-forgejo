// SPDX-License-Identifier: MIT

package integration

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	auth_model "code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/models/unittest"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/test"
	"code.gitea.io/gitea/services/f3/util"
	"code.gitea.io/gitea/services/migrations"
	"code.gitea.io/gitea/tests"

	"github.com/markbates/goth"
	"github.com/stretchr/testify/assert"
	f3_types "lab.forgefriends.org/friendlyforgeformat/gof3/config/types"
	f3_forges "lab.forgefriends.org/friendlyforgeformat/gof3/forges"
	f3_common "lab.forgefriends.org/friendlyforgeformat/gof3/forges/common"
	f3_f3 "lab.forgefriends.org/friendlyforgeformat/gof3/forges/f3"
	f3_forgejo "lab.forgefriends.org/friendlyforgeformat/gof3/forges/forgejo"
	f3_tests "lab.forgefriends.org/friendlyforgeformat/gof3/forges/tests"
	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
	f3_util "lab.forgefriends.org/friendlyforgeformat/gof3/util"
)

func TestF3_MirrorAPITOLocal(t *testing.T) {
	onGiteaRun(t, func(t *testing.T, u *url.URL) {
		AllowLocalNetworks := setting.Migrations.AllowLocalNetworks
		setting.F3.Enabled = true
		setting.Migrations.AllowLocalNetworks = true
		AppVer := setting.AppVer
		// Gitea SDK (go-sdk) need to parse the AppVer from server response, so we must set it to a valid version string.
		setting.AppVer = "1.16.0"
		defer func() {
			setting.Migrations.AllowLocalNetworks = AllowLocalNetworks
			setting.AppVer = AppVer
			migrations.Init()
		}()
		assert.NoError(t, migrations.Init())

		//
		// Step 1: create a fixture
		//
		fixtureNewF3Forge := func(t f3_tests.TestingT, logger *f3_types.Logger, user *format.User, tmpDir string) *f3_forges.ForgeRoot {
			root := f3_forges.NewForgeRoot(&f3_f3.F3{}, &f3_f3.Options{
				Options: f3_types.Options{
					Configuration: f3_types.Configuration{
						Directory: tmpDir,
					},
					Features: f3_types.AllFeatures,
					Logger:   util.ToF3Logger(nil),
				},
				Remap: true,
			})
			return root
		}
		fixture := f3_forges.NewFixture(t, f3_forges.FixtureForgeFactory{Fun: fixtureNewF3Forge, AdminRequired: false})
		fixture.NewUser(5432)
		fixture.NewMilestone()
		fixture.NewLabel()
		fixture.NewIssue()
		fixture.NewTopic()
		fixture.NewRepository()
		fixture.NewPullRequest()
		fixture.NewRelease()
		fixture.NewAsset()
		fixture.NewIssueComment(nil)
		fixture.NewPullRequestComment()
		// fixture.NewReview()
		fixture.NewIssueReaction()
		fixture.NewCommentReaction()

		//
		// Step 2: mirror F3 into Forgejo
		//
		doer, err := user_model.GetAdminUser(context.Background())
		assert.NoError(t, err)
		forgejoLocalUpload := util.ForgejoForgeRoot(f3_types.AllFeatures, doer, 0)
		upload := forgejoLocalUpload.Forge
		options := f3_common.NewMirrorOptionsRecurse()
		upload.Mirror(context.Background(), fixture.Forge, options)

		//
		// Step 3: mirror Forgejo into F3
		//
		logger := util.ToF3Logger(nil)
		f3 := f3_forges.FixtureNewF3Forge(t, logger, nil, t.TempDir())
		forgejoLocalDownload := util.ForgejoForgeRoot(f3_types.AllFeatures, doer, 0)
		download := forgejoLocalDownload.Forge
		downloadUser := download.Users.GetFromFormat(context.Background(), &format.User{UserName: fixture.UserFormat.UserName})
		downloadProject := downloadUser.Projects.GetFromFormat(context.Background(), &format.Project{Name: fixture.ProjectFormat.Name})
		options = f3_common.NewMirrorOptionsRecurse(downloadUser, downloadProject)
		f3.Forge.Mirror(context.Background(), download, options)

		//
		// Step 4: verify the fixture and F3 are equivalent
		//
		files := f3_util.Command(context.Background(), "find", f3.GetDirectory())
		assert.Contains(t, files, "/repository/git/hooks")
		assert.Contains(t, files, "/label/")
		assert.Contains(t, files, "/issue/")
		assert.Contains(t, files, "/milestone/")
		assert.Contains(t, files, "/topic/")
		assert.Contains(t, files, "/pull_request/")
		assert.Contains(t, files, "/release/")
		assert.Contains(t, files, "/asset/")
		assert.Contains(t, files, "/comment/")
		//		assert.Contains(t, files, "/review/")
		assert.Contains(t, files, "/reaction/")
		//		f3_util.Command(context.Background(), "cp", "-a", f3.GetDirectory(), "abc")
	})
}

func TestF3_MaybePromoteUser(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	//
	// OAuth2 authentication source GitLab
	//
	gitlabName := "gitlab"
	_ = addAuthSource(t, authSourcePayloadGitLabCustom(gitlabName))
	//
	// F3 authentication source matching the GitLab authentication source
	//
	f3Name := "f3"
	f3 := createF3AuthSource(t, f3Name, "http://mygitlab.eu", gitlabName)

	//
	// Create a user as if it had been previously been created by the F3
	// authentication source.
	//
	gitlabUserID := "5678"
	gitlabEmail := "gitlabuser@example.com"
	userBeforeSignIn := &user_model.User{
		Name:        "gitlabuser",
		Type:        user_model.UserTypeF3,
		LoginType:   auth_model.F3,
		LoginSource: f3.ID,
		LoginName:   gitlabUserID,
	}
	defer createUser(context.Background(), t, userBeforeSignIn)()

	//
	// A request for user information sent to Goth will return a
	// goth.User exactly matching the user created above.
	//
	defer mockCompleteUserAuth(func(res http.ResponseWriter, req *http.Request) (goth.User, error) {
		return goth.User{
			Provider: gitlabName,
			UserID:   gitlabUserID,
			Email:    gitlabEmail,
		}, nil
	})()
	req := NewRequest(t, "GET", fmt.Sprintf("/user/oauth2/%s/callback?code=XYZ&state=XYZ", gitlabName))
	resp := MakeRequest(t, req, http.StatusSeeOther)
	assert.Equal(t, "/", test.RedirectURL(resp))
	userAfterSignIn := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: userBeforeSignIn.ID})

	// both are about the same user
	assert.Equal(t, userAfterSignIn.ID, userBeforeSignIn.ID)
	// the login time was updated, proof the login succeeded
	assert.Greater(t, userAfterSignIn.LastLoginUnix, userBeforeSignIn.LastLoginUnix)
	// the login type was promoted from F3 to OAuth2
	assert.Equal(t, userBeforeSignIn.LoginType, auth_model.F3)
	assert.Equal(t, userAfterSignIn.LoginType, auth_model.OAuth2)
	// the OAuth2 email was used to set the missing user email
	assert.Equal(t, userBeforeSignIn.Email, "")
	assert.Equal(t, userAfterSignIn.Email, gitlabEmail)
}

func TestF3_UserMappingExisting(t *testing.T) {
	onGiteaRun(t, func(t *testing.T, u *url.URL) {
		AllowLocalNetworks := setting.Migrations.AllowLocalNetworks
		setting.F3.Enabled = true
		setting.Migrations.AllowLocalNetworks = true
		AppVer := setting.AppVer
		// Gitea SDK (go-sdk) need to parse the AppVer from server response, so we must set it to a valid version string.
		setting.AppVer = "1.16.0"
		defer func() {
			setting.Migrations.AllowLocalNetworks = AllowLocalNetworks
			setting.AppVer = AppVer
		}()

		log.Debug("Step 1: create a fixture in F3")
		fixtureNewF3Forge := func(t f3_tests.TestingT, logger *f3_types.Logger, user *format.User, tmpDir string) *f3_forges.ForgeRoot {
			root := f3_forges.NewForgeRoot(&f3_f3.F3{}, &f3_f3.Options{
				Options: f3_types.Options{
					Configuration: f3_types.Configuration{
						Directory: tmpDir,
					},
					Features: f3_types.AllFeatures,
					Logger:   util.ToF3Logger(nil),
				},
				Remap: true,
			})
			return root
		}
		fixture := f3_forges.NewFixture(t, f3_forges.FixtureForgeFactory{Fun: fixtureNewF3Forge, AdminRequired: false})
		userID := int64(5432)
		fixture.NewUser(userID)
		//		fixture.NewProject()

		log.Debug("Step 2: mirror F3 into Forgejo")
		//
		// OAuth2 authentication source GitLab
		//
		gitlabName := "gitlab"
		gitlab := addAuthSource(t, authSourcePayloadGitLabCustom(gitlabName))
		//
		// Create a user as if it had been previously been created by the F3
		// authentication source.
		//
		gitlabUserID := fmt.Sprintf("%d", userID)
		gitlabUser := &user_model.User{
			Name:        "gitlabuser",
			Email:       "gitlabuser@example.com",
			LoginType:   auth_model.OAuth2,
			LoginSource: gitlab.ID,
			LoginName:   gitlabUserID,
		}
		defer createUser(context.Background(), t, gitlabUser)()

		doer, err := user_model.GetAdminUser(context.Background())
		assert.NoError(t, err)
		forgejoLocal := util.ForgejoForgeRoot(f3_types.AllFeatures, doer, gitlab.ID)
		options := f3_common.NewMirrorOptionsRecurse()
		forgejoLocal.Forge.Mirror(context.Background(), fixture.Forge, options)

		log.Debug("Step 3: mirror Forgejo into F3")
		adminUsername := "user1"
		logger := util.ToF3Logger(nil)
		forgejoAPI := f3_forges.NewForgeRoot(&f3_forgejo.Forgejo{}, &f3_forgejo.Options{
			Options: f3_types.Options{
				Configuration: f3_types.Configuration{
					URL:       setting.AppURL,
					Directory: t.TempDir(),
				},
				Features: f3_types.AllFeatures,
				Logger:   logger,
			},
			AuthToken: getUserToken(t, adminUsername, auth_model.AccessTokenScopeWriteAdmin, auth_model.AccessTokenScopeAll),
		})

		f3 := f3_forges.FixtureNewF3Forge(t, logger, nil, t.TempDir())
		apiForge := forgejoAPI.Forge
		apiUser := apiForge.Users.GetFromFormat(context.Background(), &format.User{UserName: gitlabUser.Name})
		//		apiProject := apiUser.Projects.GetFromFormat(context.Background(), &format.Project{Name: fixture.ProjectFormat.Name})
		// options = f3_common.NewMirrorOptionsRecurse(apiUser, apiProject)
		options = f3_common.NewMirrorOptionsRecurse(apiUser)
		f3.Forge.Mirror(context.Background(), apiForge, options)

		//
		// Step 4: verify the fixture and F3 are equivalent
		//
		files := f3_util.Command(context.Background(), "find", f3.GetDirectory())
		assert.Contains(t, files, fmt.Sprintf("/user/%d", gitlabUser.ID))
	})
}

func TestF3_UserMappingNew(t *testing.T) {
	onGiteaRun(t, func(t *testing.T, u *url.URL) {
		AllowLocalNetworks := setting.Migrations.AllowLocalNetworks
		setting.F3.Enabled = true
		setting.Migrations.AllowLocalNetworks = true
		AppVer := setting.AppVer
		// Gitea SDK (go-sdk) need to parse the AppVer from server response, so we must set it to a valid version string.
		setting.AppVer = "1.16.0"
		defer func() {
			setting.Migrations.AllowLocalNetworks = AllowLocalNetworks
			setting.AppVer = AppVer
		}()

		log.Debug("Step 1: create a fixture in F3")
		fixtureNewF3Forge := func(t f3_tests.TestingT, logger *f3_types.Logger, user *format.User, tmpDir string) *f3_forges.ForgeRoot {
			root := f3_forges.NewForgeRoot(&f3_f3.F3{}, &f3_f3.Options{
				Options: f3_types.Options{
					Configuration: f3_types.Configuration{
						Directory: tmpDir,
					},
					Features: f3_types.AllFeatures,
					Logger:   util.ToF3Logger(nil),
				},
				Remap: true,
			})
			return root
		}
		fixture := f3_forges.NewFixture(t, f3_forges.FixtureForgeFactory{Fun: fixtureNewF3Forge, AdminRequired: false})
		userID := int64(5432)
		fixture.NewUser(userID)

		log.Debug("Step 2: mirror F3 into Forgejo")
		doer, err := user_model.GetAdminUser(context.Background())
		assert.NoError(t, err)
		forgejoLocalDestination := util.ForgejoForgeRoot(f3_types.AllFeatures, doer, 0)
		options := f3_common.NewMirrorOptionsRecurse()
		forgejoLocalDestination.Forge.Mirror(context.Background(), fixture.Forge, options)

		log.Debug("Step 3: change the Name of the user in F3 and mirror to Forgejo")
		otherusername := "otheruser"
		fixture.UserFormat.UserName = otherusername
		fixture.Forge.Users.Upsert(context.Background(), fixture.UserFormat)
		forgejoLocalDestination.Forge.Mirror(context.Background(), fixture.Forge, options)

		log.Debug("Step 4: mirror Forgejo into F3 using the changed name")
		f3 := util.F3ForgeRoot(f3_types.AllFeatures, t.TempDir())
		forgejoLocalOrigin := util.ForgejoForgeRoot(f3_types.AllFeatures, doer, 0)
		forgejoLocalOriginUser := forgejoLocalOrigin.Forge.Users.GetFromFormat(context.Background(), &format.User{UserName: otherusername})
		options = f3_common.NewMirrorOptionsRecurse(forgejoLocalOriginUser)
		f3.Forge.Mirror(context.Background(), forgejoLocalOrigin.Forge, options)

		//
		// verify the fixture and F3 are equivalent
		//
		files := f3_util.Command(context.Background(), "find", f3.GetDirectory())
		assert.Contains(t, files, fmt.Sprintf("/user/%d", forgejoLocalOriginUser.GetID()))
	})
}
