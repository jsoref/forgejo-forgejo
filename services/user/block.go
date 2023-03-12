// Copyright 2023 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT
package user

import (
	"context"

	"code.gitea.io/gitea/models/db"
	user_model "code.gitea.io/gitea/models/user"
)

// BlockUser adds a blocked user entry for userID to block blockID.
// TODO: Figure out if instance admins should be immune to blocking.
// TODO: Add more mechanism like removing blocked user as collaborator on
// repositories where the user is an owner.
func BlockUser(ctx context.Context, userID, blockID int64) error {
	if userID == blockID || user_model.IsBlocked(ctx, userID, blockID) {
		return nil
	}

	ctx, committer, err := db.TxContext(ctx)
	if err != nil {
		return err
	}
	defer committer.Close()

	// Add the blocked user entry.
	_, err = db.GetEngine(ctx).Insert(&user_model.BlockedUser{UserID: userID, BlockID: blockID})
	if err != nil {
		return err
	}

	// Unfollow the user from block's perspective.
	err = user_model.UnfollowUser(ctx, blockID, userID)
	if err != nil {
		return err
	}

	return committer.Commit()
}
