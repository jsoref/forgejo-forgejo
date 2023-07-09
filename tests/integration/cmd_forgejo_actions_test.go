// SPDX-License-Identifier: MIT

package integration

import (
	"net/url"
	"testing"

	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/test"

	"github.com/stretchr/testify/assert"
)

func Test_CmdForgejo_Actions(t *testing.T) {
	onGiteaRun(t, func(*testing.T, *url.URL) {
		defer test.MockVariable(&setting.Actions.Enabled, true)()

		var output string

		output = cmdForgejoCaptureOutput(t, []string{"forgejo-cli", "actions", "generate-runner-token"})
		assert.EqualValues(t, 40, len(output))

		output = cmdForgejoCaptureOutput(t, []string{"forgejo-cli", "actions", "generate-secret"})
		assert.EqualValues(t, 40, len(output))
	})
}
