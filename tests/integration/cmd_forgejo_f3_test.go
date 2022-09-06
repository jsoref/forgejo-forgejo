// SPDX-License-Identifier: MIT

package integration

import (
	"context"
	"testing"

	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/services/f3/driver"
	"code.gitea.io/gitea/tests"

	"github.com/stretchr/testify/assert"
	f3_forges "lab.forgefriends.org/friendlyforgeformat/gof3/forges"
	f3_util "lab.forgefriends.org/friendlyforgeformat/gof3/util"
)

func TestF3_CmdMirror_LocalForgejo(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	ctx := context.Background()
	var userID int64 = 700
	//
	// Step 1: create a fixture as an F3 archive
	//
	userID++
	fixture := f3_forges.NewFixture(t, f3_forges.FixtureF3Factory)
	fixture.NewUser(userID)
	fixture.NewIssue()
	fixture.NewRepository()

	//
	// Step 3: mirror the F3 archive to the forge
	//
	_, err := cmdForgejoCaptureOutput(t, []string{
		"forgejo", "forgejo-cli", "f3", "mirror",
		"--from-type=f3", "--from", fixture.ForgeRoot.GetDirectory(),
		"--to-type", driver.Name,
	})
	assert.NoError(t, err)
	user, err := user_model.GetUserByName(ctx, fixture.UserFormat.UserName)
	assert.NoError(t, err)
	//
	// Step 4: mirror the forge to an F3 archive
	//
	dumpDir := t.TempDir()
	_, err = cmdForgejoCaptureOutput(t, []string{
		"forgejo", "forgejo-cli", "f3", "mirror",
		"--user", user.Name, "--repository", fixture.ProjectFormat.Name,
		"--from-type", driver.Name,
		"--to-type=f3", "--to", dumpDir,
	})
	assert.NoError(t, err)

	//
	// Step 5: verify the F3 archive content
	//
	files := f3_util.Command(context.Background(), "find", dumpDir)
	assert.Contains(t, files, "/user/")
	assert.Contains(t, files, "/project/")
}
