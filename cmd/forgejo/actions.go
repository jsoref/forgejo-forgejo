// Copyright The Forgejo Authors.
// SPDX-License-Identifier: MIT

package forgejo

import (
	"context"
	"fmt"

	actions_model "code.gitea.io/gitea/models/actions"
	"code.gitea.io/gitea/modules/private"
	"code.gitea.io/gitea/modules/setting"

	"github.com/urfave/cli"
)

func CmdActions(ctx context.Context) cli.Command {
	return cli.Command{
		Name:  "actions",
		Usage: "Commands for managing Forgejo Actions",
		Subcommands: []cli.Command{
			SubcmdActionsGenerateRunnerToken(ctx),
			SubcmdActionsGenerateRunnerSecret(ctx),
		},
	}
}

func SubcmdActionsGenerateRunnerToken(ctx context.Context) cli.Command {
	return cli.Command{
		Name:   "generate-runner-token",
		Usage:  "Generate a new token for a runner to use to register with the server",
		Action: func(cliCtx *cli.Context) error { return RunGenerateActionsRunnerToken(ctx, cliCtx) },
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "scope, s",
				Value: "",
				Usage: "{owner}[/{repo}] - leave empty for a global runner",
			},
		},
	}
}

func SubcmdActionsGenerateRunnerSecret(ctx context.Context) cli.Command {
	return cli.Command{
		Name:   "generate-secret",
		Usage:  "Generate a secret suitable for input to the register subcommand",
		Action: func(cliCtx *cli.Context) error { return RunGenerateSecret(ctx, cliCtx) },
	}
}

func RunGenerateSecret(ctx context.Context, cliCtx *cli.Context) error {
	setting.MustInstalled()

	runner := actions_model.ActionRunner{}
	if err := runner.GenerateToken(); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(ContextGetStdout(ctx), "%s", runner.Token); err != nil {
		panic(err)
	}
	return nil
}

func RunGenerateActionsRunnerToken(ctx context.Context, cliCtx *cli.Context) error {
	if !ContextGetNoInstallSignals(ctx) {
		var cancel context.CancelFunc
		ctx, cancel = installSignals(ctx)
		defer cancel()
	}

	setting.MustInstalled()

	scope := cliCtx.String("scope")

	respText, extra := private.GenerateActionsRunnerToken(ctx, scope)
	if extra.HasError() {
		return handleCliResponseExtra(extra)
	}
	if _, err := fmt.Fprintf(ContextGetStdout(ctx), "%s", respText); err != nil {
		panic(err)
	}
	return nil
}
