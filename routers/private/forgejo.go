// SPDX-License-Identifier: MIT

package private

import (
	"fmt"
	"net/http"

	actions_model "code.gitea.io/gitea/models/actions"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/json"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/private"
)

func ActionsRunnerRegister(ctx *context.PrivateContext) {
	var registerRequest private.ActionsRunnerRegisterRequest
	rd := ctx.Req.Body
	defer rd.Close()

	if err := json.NewDecoder(rd).Decode(&registerRequest); err != nil {
		log.Error("%v", err)
		ctx.JSON(http.StatusInternalServerError, private.Response{
			Err: err.Error(),
		})
		return
	}

	owner, repo, err := parseScope(ctx, registerRequest.Scope)
	if err != nil {
		log.Error("%v", err)
		ctx.JSON(http.StatusInternalServerError, private.Response{
			Err: err.Error(),
		})
	}

	runner, err := actions_model.RegisterRunner(ctx, owner, repo, registerRequest.Token, registerRequest.Labels, registerRequest.Name, registerRequest.Version)
	if err != nil {
		err := fmt.Sprintf("error while registering runner: %v", err)
		log.Error("%v", err)
		ctx.JSON(http.StatusInternalServerError, private.Response{
			Err: err,
		})
		return
	}

	ctx.PlainText(http.StatusOK, runner.UUID)
}
