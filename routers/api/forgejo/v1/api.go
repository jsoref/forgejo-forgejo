// Copyright 2023 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1

import (
	gocontext "context"

	"code.gitea.io/gitea/modules/web"
)

func Routes(ctx gocontext.Context) *web.Route {
	m := web.NewRoute()
	forgejo := NewForgejo()
	m.Get("", Root)
	m.Get("/version", forgejo.GetVersion)
	return m
}
