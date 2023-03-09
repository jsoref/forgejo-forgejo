// SPDX-License-Identifier: MIT

package integration

import (
	"bytes"
	"flag"
	"io"
	"net/url"
	"os"
	"testing"

	"code.gitea.io/gitea/cmd"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func Test_CmdActions(t *testing.T) {
	onGiteaRun(t, func(*testing.T, *url.URL) {
		tests := []struct {
			name           string
			args           []string
			wantErr        bool
			expectedOutput func(string)
		}{
			{"test_registration-token-admin", []string{"actions", "--registration-token-admin"}, false, func(output string) { assert.EqualValues(t, 40, len(output), output) }},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				realStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w

				set := flag.NewFlagSet("actions", 0)
				_ = set.Parse(tt.args)
				context := cli.NewContext(&cli.App{Writer: os.Stdout}, set, nil)
				err := cmd.CmdActions.Run(context)
				if (err != nil) != tt.wantErr {
					t.Errorf("CmdActions.Run() error = %v, wantErr %v", err, tt.wantErr)
				}
				w.Close()
				var buf bytes.Buffer
				io.Copy(&buf, r)
				tt.expectedOutput(buf.String())
				os.Stdout = realStdout
			})
		}
	})
}
