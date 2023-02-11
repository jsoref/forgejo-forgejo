// Copyright 2023 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitEnvironmentVariable(t *testing.T) {
	prefixRegexp := regexp.MustCompile(prefixRegexpString + "__")
	k, v := splitEnvironmentVariable(prefixRegexp, "FORGEJO__KEY=VALUE")
	assert.Equal(t, k, "KEY")
	assert.Equal(t, v, "VALUE")
	k, v = splitEnvironmentVariable(prefixRegexp, "nothing=interesting")
	assert.Equal(t, k, "")
	assert.Equal(t, v, "")
}
