// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"net/http"
	"testing"

	auth_model "code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/tests"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	setting.AppVer = "test-version-1"
	req := NewRequest(t, "GET", "/api/v1/version")
	resp := MakeRequest(t, req, http.StatusOK)

	var version structs.ServerVersion
	DecodeJSON(t, resp, &version)
	assert.Equal(t, setting.AppVer, version.Version)

	// Verify https://codeberg.org/forgejo/forgejo/pulls/1098 is fixed
	{
		token := getUserToken(t, "user2", auth_model.AccessTokenScopeReadActivityPub)
		req := NewRequestf(t, "GET", "/api/v1/version?token=%s", token)
		resp := MakeRequest(t, req, http.StatusOK)

		var version structs.ServerVersion
		DecodeJSON(t, resp, &version)
		assert.Equal(t, setting.AppVer, version.Version)
	}
}
