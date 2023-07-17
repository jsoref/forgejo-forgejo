// Copyright 2023 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"fmt"
	"net/http"
	"testing"

	auth_model "code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/models/unittest"
	user_model "code.gitea.io/gitea/models/user"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/tests"

	"github.com/stretchr/testify/assert"
)

func TestAPIUserBlock(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	user := "user4"
	session := loginUser(t, user)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeWriteUser)

	t.Run("BlockUser", func(t *testing.T) {
		defer tests.PrintCurrentTest(t)()

		req := NewRequest(t, "PUT", fmt.Sprintf("/api/v1/user/block/user2?token=%s", token))
		MakeRequest(t, req, http.StatusNoContent)

		unittest.AssertExistsAndLoadBean(t, &user_model.BlockedUser{UserID: 4, BlockID: 2})
	})

	t.Run("ListBlocked", func(t *testing.T) {
		defer tests.PrintCurrentTest(t)()

		req := NewRequest(t, "GET", fmt.Sprintf("/api/v1/user/list_blocked?token=%s", token))
		resp := MakeRequest(t, req, http.StatusOK)

		// One user just got blocked and the other one is defined in the fixtures.
		assert.Equal(t, "2", resp.Header().Get("X-Total-Count"))

		var blockedUsers []api.BlockedUser
		DecodeJSON(t, resp, &blockedUsers)
		assert.Len(t, blockedUsers, 2)
		assert.EqualValues(t, 1, blockedUsers[0].BlockID)
		assert.EqualValues(t, 2, blockedUsers[1].BlockID)
	})

	t.Run("UnblockUser", func(t *testing.T) {
		defer tests.PrintCurrentTest(t)()

		req := NewRequest(t, "PUT", fmt.Sprintf("/api/v1/user/unblock/user2?token=%s", token))
		MakeRequest(t, req, http.StatusNoContent)

		unittest.AssertNotExistsBean(t, &user_model.BlockedUser{UserID: 4, BlockID: 2})
	})
}

func TestAPIOrgBlock(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	user := "user5"
	org := "user6"
	session := loginUser(t, user)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeWriteOrganization)

	t.Run("BlockUser", func(t *testing.T) {
		defer tests.PrintCurrentTest(t)()

		req := NewRequest(t, "PUT", fmt.Sprintf("/api/v1/orgs/%s/block/user2?token=%s", org, token))
		MakeRequest(t, req, http.StatusNoContent)

		unittest.AssertExistsAndLoadBean(t, &user_model.BlockedUser{UserID: 6, BlockID: 2})
	})

	t.Run("ListBlocked", func(t *testing.T) {
		defer tests.PrintCurrentTest(t)()

		req := NewRequest(t, "GET", fmt.Sprintf("/api/v1/orgs/%s/list_blocked?token=%s", org, token))
		resp := MakeRequest(t, req, http.StatusOK)

		assert.Equal(t, "1", resp.Header().Get("X-Total-Count"))

		var blockedUsers []api.BlockedUser
		DecodeJSON(t, resp, &blockedUsers)
		assert.Len(t, blockedUsers, 1)
		assert.EqualValues(t, 2, blockedUsers[0].BlockID)
	})

	t.Run("UnblockUser", func(t *testing.T) {
		defer tests.PrintCurrentTest(t)()

		req := NewRequest(t, "PUT", fmt.Sprintf("/api/v1/orgs/%s/unblock/user2?token=%s", org, token))
		MakeRequest(t, req, http.StatusNoContent)

		unittest.AssertNotExistsBean(t, &user_model.BlockedUser{UserID: 6, BlockID: 2})
	})
}
