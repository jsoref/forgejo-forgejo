// SPDX-License-Identifier: MIT

package driver

import (
	"context"
	"fmt"

	"code.gitea.io/gitea/models/db"
	issues_model "code.gitea.io/gitea/models/issues"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/services/convert"

	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
	"lab.forgefriends.org/friendlyforgeformat/gof3/util"
)

type Review struct {
	issues_model.Review
}

func ReviewConverter(f *issues_model.Review) *Review {
	return &Review{
		Review: *f,
	}
}

func (o Review) GetID() int64 {
	return o.ID
}

func (o Review) GetIDString() string {
	return fmt.Sprintf("%d", o.GetID())
}

func (o *Review) SetID(id int64) {
	o.ID = id
}

func (o *Review) SetIDString(id string) {
	o.SetID(util.ParseInt(id))
}

func (o *Review) IsNil() bool {
	return o.ID == 0
}

func (o *Review) Equals(other *Review) bool {
	return o.Content == other.Content
}

func (o *Review) ToFormatInterface() format.Interface {
	return o.ToFormat()
}

func (o *Review) ToFormat() *format.Review {
	comments := make([]*format.ReviewComment, 0, len(o.Comments))
	for _, comment := range o.Comments {
		comments = append(comments, &format.ReviewComment{
			Common: format.NewCommon(comment.ID),
			// InReplyTo
			Content:   comment.Content,
			TreePath:  comment.TreePath,
			DiffHunk:  convert.Patch2diff(comment.Patch),
			Patch:     comment.Patch,
			Line:      int(comment.Line),
			CommitID:  comment.CommitSHA,
			PosterID:  format.NewUserReference(comment.PosterID),
			CreatedAt: comment.CreatedUnix.AsTime(),
			UpdatedAt: comment.UpdatedUnix.AsTime(),
		})
	}

	review := format.Review{
		Common:     format.NewCommon(o.Review.ID),
		IssueIndex: o.IssueID,
		Official:   o.Review.Official,
		CommitID:   o.Review.CommitID,
		Content:    o.Review.Content,
		CreatedAt:  o.Review.CreatedUnix.AsTime(),
		State:      format.ReviewStateUnknown,
		Comments:   comments,
	}

	if o.Review.Reviewer != nil {
		review.ReviewerID = format.NewUserReference(o.Review.Reviewer.ID)
	}

	switch o.Type {
	case issues_model.ReviewTypeApprove:
		review.State = format.ReviewStateApproved
	case issues_model.ReviewTypeReject:
		review.State = format.ReviewStateChangesRequested
	case issues_model.ReviewTypeComment:
		review.State = format.ReviewStateCommented
	case issues_model.ReviewTypePending:
		review.State = format.ReviewStatePending
	case issues_model.ReviewTypeRequest:
		review.State = format.ReviewStateRequestReview
	}

	return &review
}

func (o *Review) FromFormat(review *format.Review) {
	comments := make([]*issues_model.Comment, 0, len(review.Comments))
	for _, comment := range review.Comments {
		comments = append(comments, &issues_model.Comment{
			ID:   comment.GetID(),
			Type: issues_model.CommentTypeReview,
			// InReplyTo
			CommitSHA:   comment.CommitID,
			Line:        int64(comment.Line),
			TreePath:    comment.TreePath,
			Content:     comment.Content,
			Patch:       comment.Patch,
			PosterID:    comment.PosterID.GetID(),
			CreatedUnix: timeutil.TimeStamp(comment.CreatedAt.Unix()),
			UpdatedUnix: timeutil.TimeStamp(comment.UpdatedAt.Unix()),
		})
	}
	*o = Review{
		Review: issues_model.Review{
			ID:         review.GetID(),
			ReviewerID: review.ReviewerID.GetID(),
			Reviewer: &user_model.User{
				ID: review.ReviewerID.GetID(),
			},
			IssueID:     review.IssueIndex,
			Official:    review.Official,
			CommitID:    review.CommitID,
			Content:     review.Content,
			CreatedUnix: timeutil.TimeStamp(review.CreatedAt.Unix()),
			Comments:    comments,
		},
	}

	switch review.State {
	case format.ReviewStateApproved:
		o.Type = issues_model.ReviewTypeApprove
	case format.ReviewStateChangesRequested:
		o.Type = issues_model.ReviewTypeReject
	case format.ReviewStateCommented:
		o.Type = issues_model.ReviewTypeComment
	case format.ReviewStatePending:
		o.Type = issues_model.ReviewTypePending
	case format.ReviewStateRequestReview:
		o.Type = issues_model.ReviewTypeRequest
	}
}

type ReviewProvider struct {
	BaseProvider
}

func (o *ReviewProvider) ToFormat(ctx context.Context, review *Review) *format.Review {
	return review.ToFormat()
}

func (o *ReviewProvider) FromFormat(ctx context.Context, r *format.Review) *Review {
	var review Review
	review.FromFormat(r)
	return &review
}

func (o *ReviewProvider) GetObjects(ctx context.Context, user *User, project *Project, pullRequest *PullRequest, page int) []*Review {
	reviews, err := issues_model.FindReviews(ctx, issues_model.FindReviewOptions{
		ListOptions: db.ListOptions{Page: page, PageSize: o.g.perPage},
		IssueID:     pullRequest.IssueID,
		Type:        issues_model.ReviewTypeUnknown,
	})
	if err != nil {
		panic(fmt.Errorf("error while listing reviews: %v", err))
	}

	return util.ConvertMap[*issues_model.Review, *Review](reviews, ReviewConverter)
}

func (o *ReviewProvider) ProcessObject(ctx context.Context, user *User, project *Project, pullRequest *PullRequest, review *Review) {
	if err := (&review.Review).LoadAttributes(ctx); err != nil {
		panic(err)
	}
}

func (o *ReviewProvider) Get(ctx context.Context, user *User, project *Project, pullRequest *PullRequest, exemplar *Review) *Review {
	id := exemplar.GetID()
	review, err := issues_model.GetReviewByID(ctx, id)
	if issues_model.IsErrReviewNotExist(err) {
		return &Review{}
	}
	if err != nil {
		panic(err)
	}
	if err := review.LoadAttributes(ctx); err != nil {
		panic(err)
	}
	return ReviewConverter(review)
}

func (o *ReviewProvider) Put(ctx context.Context, user *User, project *Project, pullRequest *PullRequest, review, existing *Review) *Review {
	var result *Review

	if existing == nil || existing.IsNil() {
		r := &review.Review
		r.ID = 0
		for _, comment := range r.Comments {
			comment.ID = 0
		}
		r.IssueID = pullRequest.IssueID
		if err := issues_model.InsertReviews(ctx, []*issues_model.Review{r}); err != nil {
			panic(err)
		}
		result = ReviewConverter(r)
	} else {
		var u issues_model.Review
		u.ID = existing.GetID()
		cols := make([]string, 0, 10)

		if review.Content != existing.Content {
			u.Content = review.Content
			cols = append(cols, "content")
		}
		if len(cols) > 0 {
			if _, err := db.GetEngine(ctx).ID(existing.ID).Cols(cols...).Update(u); err != nil {
				panic(err)
			}
		}
		result = existing
	}

	return o.Get(ctx, user, project, pullRequest, result)
}

func (o *ReviewProvider) Delete(ctx context.Context, user *User, project *Project, pullRequest *PullRequest, review *Review) *Review {
	r := o.Get(ctx, user, project, pullRequest, review)
	if !r.IsNil() {
		if err := issues_model.DeleteReview(ctx, &r.Review); err != nil {
			panic(err)
		}
	}
	return r
}
