// SPDX-License-Identifier: MIT

package driver

import (
	"context"
	"fmt"

	"code.gitea.io/gitea/models/db"
	issues_model "code.gitea.io/gitea/models/issues"
	user_model "code.gitea.io/gitea/models/user"

	"lab.forgefriends.org/friendlyforgeformat/gof3/forges/common"
	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
	"lab.forgefriends.org/friendlyforgeformat/gof3/util"
	"xorm.io/builder"
)

type Reaction struct {
	issues_model.Reaction
}

func ReactionConverter(f *issues_model.Reaction) *Reaction {
	return &Reaction{
		Reaction: *f,
	}
}

func (o Reaction) GetID() int64 {
	return o.ID
}

func (o Reaction) GetIDString() string {
	return fmt.Sprintf("%d", o.GetID())
}

func (o *Reaction) SetID(id int64) {
	o.ID = id
}

func (o *Reaction) SetIDString(id string) {
	o.SetID(util.ParseInt(id))
}

func (o *Reaction) IsNil() bool {
	return o.ID == 0
}

func (o *Reaction) Equals(other *Reaction) bool {
	return o.UserID == other.UserID && o.Type == other.Type
}

func (o *Reaction) ToFormatInterface() format.Interface {
	return o.ToFormat()
}

func (o *Reaction) ToFormat() *format.Reaction {
	return &format.Reaction{
		Common:  format.NewCommon(o.ID),
		UserID:  format.NewUserReference(o.User.ID),
		Content: o.Type,
	}
}

func (o *Reaction) FromFormat(reaction *format.Reaction) {
	*o = Reaction{
		Reaction: issues_model.Reaction{
			ID:     reaction.GetID(),
			UserID: reaction.UserID.GetID(),
			User: &user_model.User{
				ID: reaction.UserID.GetID(),
			},
			Type: reaction.Content,
		},
	}
}

type ReactionProvider struct {
	BaseProvider
}

func (o *ReactionProvider) ToFormat(ctx context.Context, reaction *Reaction) *format.Reaction {
	return reaction.ToFormat()
}

func (o *ReactionProvider) FromFormat(ctx context.Context, m *format.Reaction) *Reaction {
	var reaction Reaction
	reaction.FromFormat(m)
	return &reaction
}

//
// Although it would be possible to use a higher level logic instead of the database,
// as of September 2022 (1.18 dev)
// (i) models/issues/reaction.go imposes a significant overhead
// (ii) is fragile and bugous https://github.com/go-gitea/gitea/issues/20860
//

func (o *ReactionProvider) GetObjects(ctx context.Context, user *User, project *Project, parents []common.ContainerObjectInterface, page int) []*Reaction {
	cond := builder.NewCond()
	switch l := parents[len(parents)-1].(type) {
	case *Issue:
		cond = cond.And(builder.Eq{"reaction.issue_id": l.ID})
		cond = cond.And(builder.Eq{"reaction.comment_id": 0})
	case *Comment:
		cond = cond.And(builder.Eq{"reaction.comment_id": l.ID})
	default:
		panic(fmt.Errorf("unexpected type %T", parents[len(parents)-1]))
	}
	sess := db.GetEngine(ctx).Where(cond)
	if page > 0 {
		sess = db.SetSessionPagination(sess, &db.ListOptions{Page: page, PageSize: o.g.perPage})
	}
	reactions := make([]*issues_model.Reaction, 0, 10)
	if err := sess.Find(&reactions); err != nil {
		panic(err)
	}
	_, err := (issues_model.ReactionList)(reactions).LoadUsers(ctx, nil)
	if err != nil {
		panic(err)
	}
	return util.ConvertMap[*issues_model.Reaction, *Reaction](reactions, ReactionConverter)
}

func (o *ReactionProvider) ProcessObject(ctx context.Context, user *User, project *Project, parents []common.ContainerObjectInterface, reaction *Reaction) {
}

func (o *ReactionProvider) Get(ctx context.Context, user *User, project *Project, parents []common.ContainerObjectInterface, exemplar *Reaction) *Reaction {
	reaction := &Reaction{}
	has, err := db.GetEngine(ctx).ID(exemplar.GetID()).Get(&reaction.Reaction)
	if err != nil {
		panic(err)
	} else if !has {
		return &Reaction{}
	}
	if _, err := (issues_model.ReactionList{&reaction.Reaction}).LoadUsers(ctx, nil); err != nil {
		panic(err)
	}
	return reaction
}

func (o *ReactionProvider) Put(ctx context.Context, user *User, project *Project, parents []common.ContainerObjectInterface, reaction, existing *Reaction) *Reaction {
	r := &issues_model.Reaction{
		Type:   reaction.Type,
		UserID: o.g.GetDoer().ID,
	}
	switch l := parents[len(parents)-1].(type) {
	case *Issue:
		r.IssueID = l.ID
		r.CommentID = 0
	case *Comment:
		i, ok := parents[len(parents)-2].(*Issue)
		if !ok {
			panic(fmt.Errorf("unexpected type %T", parents[len(parents)-2]))
		}
		r.IssueID = i.ID
		r.CommentID = l.ID
	default:
		panic(fmt.Errorf("unexpected type %T", parents[len(parents)-1]))
	}

	ctx, committer, err := db.TxContext(ctx)
	if err != nil {
		panic(err)
	}
	defer committer.Close()

	if _, err := db.GetEngine(ctx).Insert(r); err != nil {
		panic(err)
	}

	if err := committer.Commit(); err != nil {
		panic(err)
	}
	return ReactionConverter(r)
}

func (o *ReactionProvider) Delete(ctx context.Context, user *User, project *Project, parents []common.ContainerObjectInterface, reaction *Reaction) *Reaction {
	r := o.Get(ctx, user, project, parents, reaction)
	if !r.IsNil() {
		if _, err := db.GetEngine(ctx).Delete(&reaction.Reaction); err != nil {
			panic(err)
		}
		return reaction
	}
	return r
}
