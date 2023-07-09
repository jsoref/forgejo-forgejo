// SPDX-License-Identifier: MIT

package integration

import (
	"bytes"
	"context"
	"flag"
	"io"
	"os"
	"strings"
	"testing"

	"code.gitea.io/gitea/cmd/forgejo"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func cmdForgejoCaptureOutput(t *testing.T, args []string) (string, error) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)
	set := flag.NewFlagSet("forgejo-cli", 0)
	assert.NoError(t, set.Parse(args))
	cliContext := cli.NewContext(&cli.App{Writer: w, ErrWriter: w}, set, nil)
	ctx := context.Background()
	ctx = forgejo.ContextSetNoInstallSignals(ctx, true)
	ctx = forgejo.ContextSetNoExit(ctx, true)
	ctx = forgejo.ContextSetStdout(ctx, w)
	ctx = forgejo.ContextSetStderr(ctx, w)
	if len(stdin) > 0 {
		ctx = forgejo.ContextSetStdin(ctx, strings.NewReader(strings.Join(stdin, "")))
	}
	err = forgejo.CmdForgejo(ctx).Run(cliContext)
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String(), err
}
