// SPDX-License-Identifier: MIT

package driver

import (
	"context"
	"fmt"
	"strings"

	auth_model "code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/models/db"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/util"
	user_service "code.gitea.io/gitea/services/user"

	"lab.forgefriends.org/friendlyforgeformat/gof3/forges/common"
	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
	f3_util "lab.forgefriends.org/friendlyforgeformat/gof3/util"
)

type User struct {
	user_model.User
}

func UserConverter(f *user_model.User) *User {
	return &User{
		User: *f,
	}
}

func (o User) GetID() int64 {
	return o.ID
}

func (o User) GetIDString() string {
	return fmt.Sprintf("%d", o.GetID())
}

func (o *User) SetID(id int64) {
	o.ID = id
}

func (o *User) SetIDString(id string) {
	o.SetID(f3_util.ParseInt(id))
}

func (o *User) IsNil() bool {
	return o.ID == 0
}

func (o *User) Equals(other *User) bool {
	//
	// Only compare user data if both are managed by F3 otherwise
	// they are equal if they have the same ID. Here is an example:
	//
	// * mirror from F3 to Forgejo => user jane created and assigned
	//   ID 213 & IsF3()
	// * mirror from F3 to Forgejo => user jane username in F3 is updated
	//   the username for user ID 213 in Forgejo is also updated
	// * user jane sign in with OAuth from the same source as the
	//   F3 mirror. They are promoted to IsIndividual()
	// * mirror from F3 to Forgejo => user jane username in F3 is updated
	//   the username for user ID 213 in Forgejo is **NOT** updated, it
	//   no longer is managed by F3
	//
	if !o.IsF3() || !other.IsF3() {
		return o.ID == other.ID
	}
	return (o.Name == other.Name &&
		o.FullName == other.FullName &&
		o.Email == other.Email)
}

func (o *User) ToFormatInterface() format.Interface {
	return o.ToFormat()
}

func (o *User) ToFormat() *format.User {
	return &format.User{
		Common:   format.NewCommon(o.ID),
		UserName: o.Name,
		Name:     o.FullName,
		Email:    o.Email,
		Password: o.Passwd,
	}
}

func (o *User) FromFormat(user *format.User) {
	*o = User{
		User: user_model.User{
			Type:     user_model.UserTypeF3,
			ID:       user.Index.GetID(),
			Name:     user.UserName,
			FullName: user.Name,
			Email:    user.Email,
			Passwd:   user.Password,
		},
	}
}

type UserProvider struct {
	BaseProvider
}

func getLocalMatchingRemote(ctx context.Context, authenticationSource int64, id string) *user_model.User {
	u := &user_model.User{
		LoginName:   id,
		LoginSource: authenticationSource,
		LoginType:   auth_model.OAuth2,
		Type:        user_model.UserTypeIndividual,
	}
	has, err := db.GetEngine(ctx).Get(u)
	if err != nil {
		panic(err)
	} else if !has {
		return nil
	}
	return u
}

func (o *UserProvider) GetLocalMatchingRemote(ctx context.Context, format format.Interface, parents ...common.ContainerObjectInterface) (string, bool) {
	authenticationSource := o.g.GetAuthenticationSource()
	if authenticationSource == 0 {
		return "", false
	}
	user := getLocalMatchingRemote(ctx, authenticationSource, format.GetIDString())
	if user != nil {
		o.g.GetLogger().Debug("found existing user %d with a matching authentication source for %s", user.ID, format.GetIDString())
		return fmt.Sprintf("%d", user.ID), true
	}
	o.g.GetLogger().Debug("no pre-existing local user for %s", format.GetIDString())
	return "", false
}

func (o *UserProvider) ToFormat(ctx context.Context, user *User) *format.User {
	return user.ToFormat()
}

func (o *UserProvider) FromFormat(ctx context.Context, p *format.User) *User {
	var user User
	user.FromFormat(p)
	return &user
}

func (o *UserProvider) GetObjects(ctx context.Context, page int) []*User {
	sess := db.GetEngine(ctx).In("type", user_model.UserTypeIndividual, user_model.UserTypeF3)
	if page != 0 {
		sess = db.SetSessionPagination(sess, &db.ListOptions{Page: page, PageSize: o.g.perPage})
	}
	sess = sess.Select("`user`.*")
	users := make([]*user_model.User, 0, o.g.perPage)

	if err := sess.Find(&users); err != nil {
		panic(fmt.Errorf("error while listing users: %v", err))
	}
	return f3_util.ConvertMap[*user_model.User, *User](users, UserConverter)
}

func (o *UserProvider) ProcessObject(ctx context.Context, user *User) {
}

func GetUserByName(ctx context.Context, name string) (*user_model.User, error) {
	if len(name) == 0 {
		return nil, user_model.ErrUserNotExist{Name: name}
	}
	u := &user_model.User{Name: name}
	has, err := db.GetEngine(ctx).In("type", user_model.UserTypeIndividual, user_model.UserTypeF3).Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, user_model.ErrUserNotExist{Name: name}
	}
	return u, nil
}

func (o *UserProvider) Get(ctx context.Context, exemplar *User) *User {
	o.g.GetLogger().Debug("%+v", *exemplar)
	var user *user_model.User
	var err error
	if exemplar.GetID() > 0 {
		user, err = user_model.GetUserByID(ctx, exemplar.GetID())
		o.g.GetLogger().Debug("%+v %v", user, err)
	} else if exemplar.Name != "" {
		user, err = GetUserByName(ctx, exemplar.Name)
	} else {
		panic("GetID() == 0 and UserName == \"\"")
	}
	if err != nil {
		if user_model.IsErrUserNotExist(err) {
			return &User{}
		}
		panic(fmt.Errorf("user %+v %w", *exemplar, err))
	}
	return UserConverter(user)
}

func (o *UserProvider) Put(ctx context.Context, user, existing *User) *User {
	o.g.GetLogger().Trace("begin %+v", *user)
	u := &user_model.User{
		ID:   user.GetID(),
		Type: user_model.UserTypeF3,
	}
	//
	// Get the user, if any
	//
	var has bool
	var err error
	if u.ID > 0 {
		has, err = db.GetEngine(ctx).Get(u)
		if err != nil {
			panic(err)
		}
	}
	//
	// Set user information
	//
	u.Name = user.Name
	u.LowerName = strings.ToLower(u.Name)
	u.FullName = user.FullName
	u.Email = user.Email
	if !has {
		//
		// The user does not exist, create it
		//
		o.g.GetLogger().Trace("creating %+v", *u)
		u.ID = 0
		u.Passwd = user.Passwd
		overwriteDefault := &user_model.CreateUserOverwriteOptions{
			IsActive: util.OptionalBoolTrue,
		}
		err := user_model.CreateUser(ctx, u, overwriteDefault)
		if err != nil {
			panic(err)
		}
	} else {
		//
		// The user already exists, update it
		//
		o.g.GetLogger().Trace("updating %+v", *u)
		if err := user_model.UpdateUserCols(ctx, u, "name", "lower_name", "email", "full_name"); err != nil {
			panic(err)
		}
	}
	r := o.Get(ctx, UserConverter(u))
	o.g.GetLogger().Trace("finish %+v", r.User)
	return r
}

func (o *UserProvider) Delete(ctx context.Context, user *User) *User {
	u := o.Get(ctx, user)
	if !u.IsNil() {
		if err := user_service.DeleteUser(ctx, &user.User, true); err != nil {
			panic(err)
		}
	}
	return u
}
