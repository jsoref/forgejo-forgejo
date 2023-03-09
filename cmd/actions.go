// SPDX-License-Identifier: MIT

package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"

	actions_model "code.gitea.io/gitea/models/actions"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/util"

	"github.com/urfave/cli"
)

var (
	// CmdActions represents the available actions sub-commands.
	CmdActions = cli.Command{
		Name:        "actions",
		Usage:       "",
		Description: "Commands for managing Gitea Actions",
		Subcommands: []cli.Command{
			subcmdActionsGenRunnerToken,
		},
	}

	subcmdActionsGenRunnerToken = cli.Command{
		Name:    "generate-runner-token",
		Usage:   "Generate a new token for a runner to use to register with the server",
		Action:  runGenerateActionsRunnerToken,
		Aliases: []string{"grt"},
		Flags:   []cli.Flag{},
	}
)

func maybeInitDB(stdCtx context.Context) error {
	if setting.Database.Type == "" {
		if err := initDB(stdCtx); err != nil {
			return err
		}
	}
	return nil
}

func runGenerateActionsRunnerToken(ctx *cli.Context) error {
	stdCtx := context.Background()

	if err := maybeInitDB(stdCtx); err != nil {
		log.Fatalf("maybeInitDB %v", err)
	}

	// ownid=0,repo_id=0,means this token is used for global
	return runActionsRegistrationToken(stdCtx, 0, 0)
}

func runActionsRegistrationToken(stdCtx context.Context, ownerID, repoID int64) error {
	var token *actions_model.ActionRunnerToken
	token, err := actions_model.GetUnactivatedRunnerToken(stdCtx, ownerID, repoID)
	if errors.Is(err, util.ErrNotExist) {
		token, err = actions_model.NewRunnerToken(stdCtx, ownerID, repoID)
		if err != nil {
			log.Fatalf("CreateRunnerToken %v", err)
		}
	} else if err != nil {
		log.Fatalf("GetUnactivatedRunnerToken %v", err)
	}
	fmt.Print(token.Token)
	return nil
}
