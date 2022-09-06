// SPDX-License-Identifier: MIT

package forgejo

import (
	"context"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"

	_ "code.gitea.io/gitea/services/f3/driver" // register the driver

	"github.com/urfave/cli/v2"
	f3_cmd "lab.forgefriends.org/friendlyforgeformat/gof3/cmd"
	f3_types "lab.forgefriends.org/friendlyforgeformat/gof3/config/types"
)

func F3Logger() *f3_types.Logger {
	messenger := func(message string, args ...interface{}) {
		log.Info("Message: "+message, args...)
	}
	return &f3_types.Logger{
		Message:  f3_types.LoggerFun(messenger),
		Trace:    log.Trace,
		Debug:    log.Debug,
		Info:     log.Info,
		Warn:     log.Warn,
		Error:    log.Error,
		Critical: log.Critical,
		Fatal:    log.Fatal,
	}
}

func CmdF3(ctx context.Context) *cli.Command {
	ctx = f3_types.ContextSetLogger(ctx, F3Logger())
	return &cli.Command{
		Name:  "f3",
		Usage: "F3",
		Subcommands: []*cli.Command{
			SubcmdF3Mirror(ctx),
		},
	}
}

func SubcmdF3Mirror(ctx context.Context) *cli.Command {
	mirrorCmd := f3_cmd.CreateCmdMirror(ctx)
	mirrorCmd.Before = prepareWorkPathAndCustomConf(ctx)
	f3Action := mirrorCmd.Action
	mirrorCmd.Action = func(c *cli.Context) error { return runMirror(ctx, c, f3Action) }
	return mirrorCmd
}

func runMirror(ctx context.Context, c *cli.Context, action cli.ActionFunc) error {
	var cancel context.CancelFunc
	if !ContextGetNoInit(ctx) {
		ctx, cancel = installSignals(ctx)
		defer cancel()

		if err := initDB(ctx); err != nil {
			return err
		}

		if err := git.InitSimple(ctx); err != nil {
			return err
		}
		if err := models.Init(ctx); err != nil {
			return err
		}
	}

	return action(c)
}
