// Copyright 2023 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package user

import (
	"context"
	"errors"

	"code.gitea.io/gitea/models/db"
	"code.gitea.io/gitea/modules/timeutil"
)

// ErrBlockedByUser defines an error stating that the user is not allowed to perform the action because they are blocked.
var ErrBlockedByUser = errors.New("user is blocked by the poster or repository owner")

// BlockedUser represents a blocked user entry.
type BlockedUser struct {
	ID int64 `xorm:"pk autoincr"`
	// UID of the one who got blocked.
	BlockID int64 `xorm:"index"`
	// UID of the one who did the block action.
	UserID int64 `xorm:"index"`

	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

// TableName provides the real table name
func (*BlockedUser) TableName() string {
	return "forgejo_blocked_user"
}

func init() {
	db.RegisterModel(new(BlockedUser))
}

// IsBlocked returns if userID has blocked blockID.
func IsBlocked(ctx context.Context, userID, blockID int64) bool {
	has, _ := db.GetEngine(ctx).Exist(&BlockedUser{UserID: userID, BlockID: blockID})
	return has
}

// IsBlockedMultiple returns if one of the userIDs has blocked blockID.
func IsBlockedMultiple(ctx context.Context, userIDs []int64, blockID int64) bool {
	has, _ := db.GetEngine(ctx).In("user_id", userIDs).Exist(&BlockedUser{BlockID: blockID})
	return has
}

// UnblockUser removes the blocked user entry.
func UnblockUser(ctx context.Context, userID, blockID int64) error {
	_, err := db.GetEngine(ctx).Delete(&BlockedUser{UserID: userID, BlockID: blockID})
	return err
}

// ListBlockedUsers returns the users that the user has blocked.
func ListBlockedUsers(ctx context.Context, userID int64) ([]*User, error) {
	users := make([]*User, 0, 8)
	err := db.GetEngine(ctx).
		Select("`user`.*").
		Join("INNER", "forgejo_blocked_user", "`user`.id=`forgejo_blocked_user`.block_id").
		Where("`forgejo_blocked_user`.user_id=?", userID).
		Find(&users)

	return users, err
}

// ListBlockedByUsersID returns the ids of the users that blocked the user.
func ListBlockedByUsersID(ctx context.Context, userID int64) ([]int64, error) {
	users := make([]int64, 0, 8)
	err := db.GetEngine(ctx).
		Table("user").
		Select("`user`.id").
		Join("INNER", "forgejo_blocked_user", "`user`.id=`forgejo_blocked_user`.user_id").
		Where("`forgejo_blocked_user`.block_id=?", userID).
		Find(&users)

	return users, err
}
