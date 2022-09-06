// SPDX-FileCopyrightText: Copyright the Forgejo contributors
// SPDX-License-Identifier: MIT

package f3

import (
	"code.gitea.io/gitea/models/auth"
	"code.gitea.io/gitea/modules/json"
)

type Source struct {
	URL            string
	MatchingSource string

	// reference to the authSource
	authSource *auth.Source
}

func (source *Source) FromDB(bs []byte) error {
	return json.UnmarshalHandleDoubleEncode(bs, &source)
}

func (source *Source) ToDB() ([]byte, error) {
	return json.Marshal(source)
}

func (source *Source) SetAuthSource(authSource *auth.Source) {
	source.authSource = authSource
}

func init() {
	auth.RegisterTypeConfig(auth.F3, &Source{})
}
