// SPDX-License-Identifier: MIT

package util

import (
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/log"
	base "code.gitea.io/gitea/modules/migration"
	"code.gitea.io/gitea/services/f3/driver"

	f3_types "lab.forgefriends.org/friendlyforgeformat/gof3/config/types"
	f3_forges "lab.forgefriends.org/friendlyforgeformat/gof3/forges"
	"lab.forgefriends.org/friendlyforgeformat/gof3/forges/f3"
)

func ToF3Logger(messenger base.Messenger) *f3_types.Logger {
	if messenger == nil {
		messenger = func(message string, args ...interface{}) {
			log.Info("Message: "+message, args...)
		}
	}
	return &f3_types.Logger{
		Message:  f3_types.LoggerFun(messenger),
		Trace:    log.Trace,
		Debug:    log.Debug,
		Info:     log.Info,
		Warn:     log.Warn,
		Error:    log.Error,
		Critical: log.Critical,
		Fatal:    log.Fatal,
	}
}

func ForgejoForgeRoot(features f3_types.Features, doer *user_model.User, authenticationSource int64) *f3_forges.ForgeRoot {
	forgeRoot := f3_forges.NewForgeRoot(&driver.Forgejo{}, &driver.Options{
		Options: f3_types.Options{
			Features: features,
			Logger:   ToF3Logger(nil),
		},
		Doer:                 doer,
		AuthenticationSource: authenticationSource,
	})
	return forgeRoot
}

func F3ForgeRoot(features f3_types.Features, directory string) *f3_forges.ForgeRoot {
	forgeRoot := f3_forges.NewForgeRoot(&f3.F3{}, &f3.Options{
		Options: f3_types.Options{
			Configuration: f3_types.Configuration{
				Directory: directory,
			},
			Features: features,
			Logger:   ToF3Logger(nil),
		},
		Remap: true,
	})
	return forgeRoot
}
