// SPDX-License-Identifier: MIT

package driver

import (
	"testing"

	issue_model "code.gitea.io/gitea/models/issues"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/timeutil"

	"lab.forgefriends.org/friendlyforgeformat/gof3/forges/tests"
	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
)

func TestF3Driver_IssueFormat(t *testing.T) {
	now := timeutil.TimeStampNow()
	updated := now.Add(1)
	closed := now.Add(2)
	issue := Issue{
		Issue: issue_model.Issue{
			Title:    "title",
			Index:    123,
			PosterID: 11111,
			Poster: &user_model.User{
				ID: 11111,
			},
			Content: "content",
			Milestone: &issue_model.Milestone{
				Name: "milestone1",
			},
			IsClosed:    true,
			CreatedUnix: now,
			UpdatedUnix: updated,
			ClosedUnix:  closed,
			IsLocked:    false,
			Labels: []*issue_model.Label{
				{
					Name: "label1",
				},
			},
			Assignees: []*user_model.User{
				{
					Name: "assignee1",
				},
			},
		},
	}
	tests.ToFromFormat[Issue, format.Issue, *Issue, *format.Issue](t, &issue)
}
