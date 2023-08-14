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

	t.Run("Follow", func(t *testing.T) {
		defer user_model.UnblockUser(db.DefaultContext, doer.ID, blockedUser.ID)

		// Follow each other.
		assert.NoError(t, user_model.FollowUser(db.DefaultContext, doer.ID, blockedUser.ID))
		assert.NoError(t, user_model.FollowUser(db.DefaultContext, blockedUser.ID, doer.ID))

		assert.NoError(t, BlockUser(db.DefaultContext, doer.ID, blockedUser.ID))

		// Ensure they aren't following each other anymore.
		assert.False(t, user_model.IsFollowing(db.DefaultContext, doer.ID, blockedUser.ID))
		assert.False(t, user_model.IsFollowing(db.DefaultContext, blockedUser.ID, doer.ID))
	})

	t.Run("Watch", func(t *testing.T) {
		defer user_model.UnblockUser(db.DefaultContext, doer.ID, blockedUser.ID)

		// Blocked user watch repository of doer.
		repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{OwnerID: doer.ID})
		assert.NoError(t, repo_model.WatchRepo(db.DefaultContext, blockedUser.ID, repo.ID, true))

		assert.NoError(t, BlockUser(db.DefaultContext, doer.ID, blockedUser.ID))

		// Ensure blocked user isn't following doer's repository.
		assert.False(t, repo_model.IsWatching(db.DefaultContext, blockedUser.ID, repo.ID))
	})

	t.Run("Collaboration", func(t *testing.T) {
		doer := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 16})
		blockedUser := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 18})
		repo1 := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 22, OwnerID: doer.ID})
		repo2 := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 21, OwnerID: doer.ID})
		defer user_model.UnblockUser(db.DefaultContext, doer.ID, blockedUser.ID)

		isBlockedUserCollab := func(repo *repo_model.Repository) bool {
			isCollaborator, err := repo_model.IsCollaborator(db.DefaultContext, repo.ID, blockedUser.ID)
			assert.NoError(t, err)
			return isCollaborator
		}

		assert.True(t, isBlockedUserCollab(repo1))
		assert.True(t, isBlockedUserCollab(repo2))

		assert.NoError(t, BlockUser(db.DefaultContext, doer.ID, blockedUser.ID))

		assert.False(t, isBlockedUserCollab(repo1))
		assert.False(t, isBlockedUserCollab(repo2))
	})
}
