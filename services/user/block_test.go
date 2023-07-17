// Copyright 2023 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package user

import (
	"testing"

	"code.gitea.io/gitea/models/db"
	repo_model "code.gitea.io/gitea/models/repo"
	"code.gitea.io/gitea/models/unittest"
	user_model "code.gitea.io/gitea/models/user"

	"github.com/stretchr/testify/assert"
)

// TestBlockUser will ensure that when you block a user, certain actions have
// been taken, like unfollowing each other etc.
func TestBlockUser(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	doer := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 5})
	blockedUser := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 1})

	// Follow each other.
	assert.NoError(t, user_model.FollowUser(db.DefaultContext, doer.ID, blockedUser.ID))
	assert.NoError(t, user_model.FollowUser(db.DefaultContext, blockedUser.ID, doer.ID))

	// Blocked user watch repository of doer.
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{OwnerID: doer.ID})
	assert.NoError(t, repo_model.WatchRepo(db.DefaultContext, blockedUser.ID, repo.ID, true))

	assert.NoError(t, BlockUser(db.DefaultContext, doer.ID, blockedUser.ID))

	// Ensure they aren't following each other anymore.
	assert.False(t, user_model.IsFollowing(doer.ID, blockedUser.ID))
	assert.False(t, user_model.IsFollowing(blockedUser.ID, doer.ID))

	// Ensure blocked user isn't following doer's repository.
	assert.False(t, repo_model.IsWatching(blockedUser.ID, repo.ID))
}
