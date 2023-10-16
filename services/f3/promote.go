// SPDX-FileCopyrightText: Copyright the Forgejo contributors
// SPDX-License-Identifier: MIT

package f3

import (
	"context"

	auth_model "code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/models/db"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/log"
	f3_source "code.gitea.io/gitea/services/auth/source/f3"
	"code.gitea.io/gitea/services/auth/source/oauth2"
)

func getUserByLoginName(ctx context.Context, name string) (*user_model.User, error) {
	if len(name) == 0 {
		return nil, user_model.ErrUserNotExist{Name: name}
	}
	u := &user_model.User{LoginName: name, LoginType: auth_model.F3, Type: user_model.UserTypeF3}
	has, err := db.GetEngine(ctx).Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, user_model.ErrUserNotExist{Name: name}
	}
	return u, nil
}

// The user created by F3 has:
//
//	Type        UserTypeF3
//	LogingType  F3
//	LoginName   set to the unique identifier of the originating forge
//	LoginSource set to the F3 source that can be matched against a OAuth2 source
//
// If the source from which an authentification happens is OAuth2, an existing
// F3 user will be promoted to an OAuth2 user provided:
//
//	user.LoginName is the same as goth.UserID (argument loginName)
//	user.LoginSource has a MatchingSource equals to the name of the OAuth2 provider
//
// Once promoted, the user will be logged in without further interaction from the
// user and will own all repositories, issues, etc. associated with it.
func MaybePromoteF3User(ctx context.Context, source *auth_model.Source, loginName, email string) error {
	user, err := getF3UserToPromote(ctx, source, loginName, email)
	if err != nil {
		return err
	}
	if user != nil {
		promote := &user_model.User{
			ID:          user.ID,
			Type:        user_model.UserTypeIndividual,
			Email:       email,
			LoginSource: source.ID,
			LoginType:   source.Type,
		}
		log.Debug("promote user %v: LoginName %v => %v, LoginSource %v => %v, LoginType %v => %v, Email %v => %v", user.ID, user.LoginName, promote.LoginName, user.LoginSource, promote.LoginSource, user.LoginType, promote.LoginType, user.Email, promote.Email)
		return user_model.UpdateUser(ctx, promote, true, "type", "email", "login_source", "login_type")
	}
	return nil
}

func getF3UserToPromote(ctx context.Context, source *auth_model.Source, loginName, email string) (*user_model.User, error) {
	if !source.IsOAuth2() {
		log.Debug("getF3UserToPromote: source %v is not OAuth2", source)
		return nil, nil
	}
	oauth2Source, ok := source.Cfg.(*oauth2.Source)
	if !ok {
		log.Error("getF3UserToPromote: source claims to be OAuth2 but really is %v", oauth2Source)
		return nil, nil
	}

	u, err := getUserByLoginName(ctx, loginName)
	if err != nil {
		if user_model.IsErrUserNotExist(err) {
			log.Debug("getF3UserToPromote: no user with LoginType F3 and LoginName '%s'", loginName)
			return nil, nil
		}
		return nil, err
	}

	if !u.IsF3() {
		log.Debug("getF3UserToPromote: user %v is not a managed by F3", u)
		return nil, nil
	}

	if u.Email != "" {
		log.Debug("getF3UserToPromote: the user email is already set to '%s'", u.Email)
		return nil, nil
	}

	userSource, err := auth_model.GetSourceByID(ctx, u.LoginSource)
	if err != nil {
		if auth_model.IsErrSourceNotExist(err) {
			log.Error("getF3UserToPromote: source id = %v for user %v not found %v", u.LoginSource, u.ID, err)
			return nil, nil
		}
		return nil, err
	}
	f3Source, ok := userSource.Cfg.(*f3_source.Source)
	if !ok {
		log.Error("getF3UserToPromote: expected an F3 source but got %T %v", userSource, userSource)
		return nil, nil
	}

	if oauth2Source.Provider != f3Source.MatchingSource {
		log.Debug("getF3UserToPromote: skip OAuth2 source %s because it is different from %s which is the expected match for the F3 source %s", oauth2Source.Provider, f3Source.MatchingSource, f3Source.URL)
		return nil, nil
	}

	return u, nil
}
