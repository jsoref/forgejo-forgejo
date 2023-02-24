// SPDX-License-Identifier: MIT

package context

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOtpHeader(t *testing.T) {
	header := http.Header{}
	assert.EqualValues(t, "", getOtpHeader(header))
	// Gitea
	giteaOtp := "123456"
	header.Set("X-Gitea-OTP", giteaOtp)
	assert.EqualValues(t, giteaOtp, getOtpHeader(header))
	// Forgejo has precedence
	forgejoOtp := "abcdef"
	header.Set("X-Forgejo-OTP", forgejoOtp)
	assert.EqualValues(t, forgejoOtp, getOtpHeader(header))
}
