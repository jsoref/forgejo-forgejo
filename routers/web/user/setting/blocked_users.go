// Copyright 2023 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"net/http"

	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/setting"
)

const (
	tplSettingsBlockedUsers base.TplName = "user/settings/blocked_users"
)

// BlockedUsers render the blocked users list page.
func BlockedUsers(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("settings.blocked_users")
	ctx.Data["PageIsBlockedUsers"] = true
	ctx.Data["BaseLink"] = setting.AppSubURL + "/user/settings/blocked_users"
	ctx.Data["BaseLinkNew"] = setting.AppSubURL + "/user/settings/blocked_users"

	blockedUsers, err := user_model.ListBlockedUsers(ctx, ctx.Doer.ID)
	if err != nil {
		ctx.ServerError("ListBlockedUsers", err)
		return
	}

	ctx.Data["BlockedUsers"] = blockedUsers
	ctx.HTML(http.StatusOK, tplSettingsBlockedUsers)
}
