// SPDX-License-Identifier: MIT

package driver

import (
	"context"
	"fmt"
	"time"

	"code.gitea.io/gitea/models/db"
	issues_model "code.gitea.io/gitea/models/issues"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/timeutil"

	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
	"lab.forgefriends.org/friendlyforgeformat/gof3/util"
)

type Milestone struct {
	issues_model.Milestone
}

func MilestoneConverter(f *issues_model.Milestone) *Milestone {
	return &Milestone{
		Milestone: *f,
	}
}

func (o Milestone) GetID() int64 {
	return o.ID
}

func (o Milestone) GetIDString() string {
	return fmt.Sprintf("%d", o.GetID())
}

func (o Milestone) GetName() string {
	return o.Name
}

func (o *Milestone) SetID(id int64) {
	o.ID = id
}

func (o *Milestone) SetIDString(id string) {
	o.SetID(util.ParseInt(id))
}

func (o *Milestone) IsNil() bool {
	return o.ID == 0
}

func (o *Milestone) Equals(other *Milestone) bool {
	return o.Name == other.Name
}

func (o *Milestone) ToFormatInterface() format.Interface {
	return o.ToFormat()
}

func (o *Milestone) ToFormat() *format.Milestone {
	milestone := &format.Milestone{
		Common:      format.NewCommon(o.ID),
		Title:       o.Name,
		Description: o.Content,
		Created:     o.CreatedUnix.AsTime(),
		Updated:     o.UpdatedUnix.AsTimePtr(),
		State:       string(o.State()),
	}
	if o.IsClosed {
		milestone.Closed = o.ClosedDateUnix.AsTimePtr()
	}
	if o.DeadlineUnix.Year() < 9999 {
		milestone.Deadline = o.DeadlineUnix.AsTimePtr()
	}
	return milestone
}

func (o *Milestone) FromFormat(milestone *format.Milestone) {
	var deadline timeutil.TimeStamp
	if milestone.Deadline != nil {
		deadline = timeutil.TimeStamp(milestone.Deadline.Unix())
	}
	if deadline == 0 {
		deadline = timeutil.TimeStamp(time.Date(9999, 1, 1, 0, 0, 0, 0, setting.DefaultUILocation).Unix())
	}

	var closed timeutil.TimeStamp
	if milestone.Closed != nil {
		closed = timeutil.TimeStamp(milestone.Closed.Unix())
	}

	if milestone.Created.IsZero() {
		if milestone.Updated != nil {
			milestone.Created = *milestone.Updated
		} else if milestone.Deadline != nil {
			milestone.Created = *milestone.Deadline
		} else {
			milestone.Created = time.Now()
		}
	}
	if milestone.Updated == nil || milestone.Updated.IsZero() {
		milestone.Updated = &milestone.Created
	}

	*o = Milestone{
		issues_model.Milestone{
			ID:             milestone.GetID(),
			Name:           milestone.Title,
			Content:        milestone.Description,
			IsClosed:       milestone.State == "closed",
			CreatedUnix:    timeutil.TimeStamp(milestone.Created.Unix()),
			UpdatedUnix:    timeutil.TimeStamp(milestone.Updated.Unix()),
			ClosedDateUnix: closed,
			DeadlineUnix:   deadline,
		},
	}
}

type MilestoneProvider struct {
	BaseProviderWithProjectProvider
}

func (o *MilestoneProvider) ToFormat(ctx context.Context, milestone *Milestone) *format.Milestone {
	return milestone.ToFormat()
}

func (o *MilestoneProvider) FromFormat(ctx context.Context, m *format.Milestone) *Milestone {
	var milestone Milestone
	milestone.FromFormat(m)
	return &milestone
}

func (o *MilestoneProvider) GetObjects(ctx context.Context, user *User, project *Project, page int) []*Milestone {
	milestones, _, err := issues_model.GetMilestones(issues_model.GetMilestonesOption{
		ListOptions: db.ListOptions{Page: page, PageSize: o.g.perPage},
		RepoID:      project.GetID(),
		State:       api.StateAll,
	})
	if err != nil {
		panic(fmt.Errorf("error while listing milestones: %v", err))
	}

	r := util.ConvertMap[*issues_model.Milestone, *Milestone](([]*issues_model.Milestone)(milestones), MilestoneConverter)
	if o.project != nil {
		o.project.milestones = util.NewNameIDMap[*Milestone](r)
	}
	return r
}

func (o *MilestoneProvider) ProcessObject(ctx context.Context, user *User, project *Project, milestone *Milestone) {
}

func (o *MilestoneProvider) Get(ctx context.Context, user *User, project *Project, exemplar *Milestone) *Milestone {
	id := exemplar.GetID()
	milestone, err := issues_model.GetMilestoneByRepoID(ctx, project.GetID(), id)
	if issues_model.IsErrMilestoneNotExist(err) {
		return &Milestone{}
	}
	if err != nil {
		panic(err)
	}
	return MilestoneConverter(milestone)
}

func (o *MilestoneProvider) Put(ctx context.Context, user *User, project *Project, milestone, existing *Milestone) *Milestone {
	m := milestone.Milestone
	m.RepoID = project.GetID()

	var result *Milestone

	if existing == nil || existing.IsNil() {
		if err := issues_model.NewMilestone(ctx, &m); err != nil {
			panic(err)
		}
		result = MilestoneConverter(&m)
	} else {
		var u issues_model.Milestone
		u.ID = existing.GetID()
		cols := make([]string, 0, 10)

		if m.Name != existing.Name {
			u.Name = m.Name
			cols = append(cols, "name")
		}

		if len(cols) > 0 {
			if _, err := db.GetEngine(ctx).ID(existing.ID).Cols(cols...).Update(u); err != nil {
				panic(err)
			}
		}

		result = existing
	}

	return o.Get(ctx, user, project, result)
}

func (o *MilestoneProvider) Delete(ctx context.Context, user *User, project *Project, milestone *Milestone) *Milestone {
	m := o.Get(ctx, user, project, milestone)
	if !m.IsNil() {
		if err := issues_model.DeleteMilestoneByRepoID(ctx, project.GetID(), m.GetID()); err != nil {
			panic(err)
		}
	}
	return m
}
