// SPDX-License-Identifier: MIT

package driver

import (
	"testing"

	user_model "code.gitea.io/gitea/models/user"

	"lab.forgefriends.org/friendlyforgeformat/gof3/forges/tests"
	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
)

func TestF3Driver_UserFormat(t *testing.T) {
	user := User{
		User: user_model.User{
			ID:       1234,
			Type:     user_model.UserTypeF3,
			Name:     "username",
			FullName: "User Name",
			Email:    "username@example.com",
		},
	}
	tests.ToFromFormat[User, format.User, *User, *format.User](t, &user)
}
