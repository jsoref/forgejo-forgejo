// Copyright 2023 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"testing"

	"code.gitea.io/gitea/models/db"
	issue_model "code.gitea.io/gitea/models/issues"
	repo_model "code.gitea.io/gitea/models/repo"
	"code.gitea.io/gitea/models/unittest"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/translation"
	"code.gitea.io/gitea/tests"

	"github.com/stretchr/testify/assert"
)

func BlockUser(t *testing.T, doer, blockedUser *user_model.User) {
	t.Helper()

	unittest.AssertNotExistsBean(t, &user_model.BlockedUser{BlockID: blockedUser.ID, UserID: doer.ID})

	session := loginUser(t, doer.Name)
	req := NewRequestWithValues(t, "POST", "/"+blockedUser.Name, map[string]string{
		"_csrf":  GetCSRF(t, session, "/"+blockedUser.Name),
		"action": "block",
	})
	resp := session.MakeRequest(t, req, http.StatusOK)

	type redirect struct {
		Redirect string `json:"redirect"`
	}

	var respBody redirect
	DecodeJSON(t, resp, &respBody)
	assert.EqualValues(t, "/"+blockedUser.Name, respBody.Redirect)
	assert.EqualValues(t, true, unittest.BeanExists(t, &user_model.BlockedUser{BlockID: blockedUser.ID, UserID: doer.ID}))
}

func TestBlockUser(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	doer := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 8})
	blockedUser := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 1})
	BlockUser(t, doer, blockedUser)

	// Unblock user.
	session := loginUser(t, doer.Name)
	req := NewRequestWithValues(t, "POST", "/"+blockedUser.Name, map[string]string{
		"_csrf":  GetCSRF(t, session, "/"+blockedUser.Name),
		"action": "unblock",
	})
	resp := session.MakeRequest(t, req, http.StatusSeeOther)

	loc := resp.Header().Get("Location")
	assert.EqualValues(t, "/"+blockedUser.Name, loc)
	unittest.AssertNotExistsBean(t, &user_model.BlockedUser{BlockID: blockedUser.ID, UserID: doer.ID})
}

func TestBlockIssueCreation(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	doer := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2})
	blockedUser := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 1})
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 2, OwnerID: doer.ID})
	BlockUser(t, doer, blockedUser)

	session := loginUser(t, blockedUser.Name)
	req := NewRequest(t, "GET", "/"+repo.OwnerName+"/"+repo.Name+"/issues/new")
	resp := session.MakeRequest(t, req, http.StatusOK)

	htmlDoc := NewHTMLParser(t, resp.Body)
	link, exists := htmlDoc.doc.Find("form.ui.form").Attr("action")
	assert.True(t, exists)
	req = NewRequestWithValues(t, "POST", link, map[string]string{
		"_csrf":   htmlDoc.GetCSRF(),
		"title":   "Title",
		"content": "Hello!",
	})

	resp = session.MakeRequest(t, req, http.StatusOK)
	htmlDoc = NewHTMLParser(t, resp.Body)
	assert.Contains(t,
		htmlDoc.doc.Find(".ui.negative.message").Text(),
		translation.NewLocale("en-US").Tr("repo.issues.blocked_by_user"),
	)
}

func TestBlockIssueReaction(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	doer := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2})
	blockedUser := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 1})
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 2})
	issue := unittest.AssertExistsAndLoadBean(t, &issue_model.Issue{ID: 4, PosterID: doer.ID, RepoID: repo.ID})
	issueURL := fmt.Sprintf("/%s/%s/issues/%d", url.PathEscape(repo.OwnerName), url.PathEscape(repo.Name), issue.Index)

	BlockUser(t, doer, blockedUser)

	session := loginUser(t, blockedUser.Name)
	req := NewRequest(t, "GET", issueURL)
	resp := session.MakeRequest(t, req, http.StatusOK)
	htmlDoc := NewHTMLParser(t, resp.Body)

	req = NewRequestWithValues(t, "POST", path.Join(issueURL, "/reactions/react"), map[string]string{
		"_csrf":   htmlDoc.GetCSRF(),
		"content": "eyes",
	})
	resp = session.MakeRequest(t, req, http.StatusOK)
	type reactionResponse struct {
		Empty bool `json:"empty"`
	}

	var respBody reactionResponse
	DecodeJSON(t, resp, &respBody)

	assert.EqualValues(t, true, respBody.Empty)
}

func TestBlockCommentReaction(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	doer := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 5})
	blockedUser := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 1})
	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 1})
	issue := unittest.AssertExistsAndLoadBean(t, &issue_model.Issue{ID: 1, RepoID: repo.ID})
	comment := unittest.AssertExistsAndLoadBean(t, &issue_model.Comment{ID: 3, PosterID: doer.ID, IssueID: issue.ID})
	_ = comment.LoadIssue(db.DefaultContext)
	issueURL := fmt.Sprintf("/%s/%s/issues/%d", url.PathEscape(repo.OwnerName), url.PathEscape(repo.Name), issue.Index)

	BlockUser(t, doer, blockedUser)

	session := loginUser(t, blockedUser.Name)
	req := NewRequest(t, "GET", issueURL)
	resp := session.MakeRequest(t, req, http.StatusOK)
	htmlDoc := NewHTMLParser(t, resp.Body)

	req = NewRequestWithValues(t, "POST", path.Join(repo.Link(), "/comments/", strconv.FormatInt(comment.ID, 10), "/reactions/react"), map[string]string{
		"_csrf":   htmlDoc.GetCSRF(),
		"content": "eyes",
	})
	resp = session.MakeRequest(t, req, http.StatusOK)
	type reactionResponse struct {
		Empty bool `json:"empty"`
	}

	var respBody reactionResponse
	DecodeJSON(t, resp, &respBody)

	assert.EqualValues(t, true, respBody.Empty)
}
