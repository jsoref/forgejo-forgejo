// SPDX-License-Identifier: MIT

package driver

import (
	"context"
	"fmt"

	"code.gitea.io/gitea/models/db"
	repo_model "code.gitea.io/gitea/models/repo"

	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
	"lab.forgefriends.org/friendlyforgeformat/gof3/util"
)

type Topic struct {
	repo_model.Topic
}

func TopicConverter(f *repo_model.Topic) *Topic {
	return &Topic{
		Topic: *f,
	}
}

func (o Topic) GetID() int64 {
	return o.ID
}

func (o Topic) GetIDString() string {
	return fmt.Sprintf("%d", o.GetID())
}

func (o *Topic) SetID(id int64) {
	o.ID = id
}

func (o *Topic) SetIDString(id string) {
	o.SetID(util.ParseInt(id))
}

func (o *Topic) IsNil() bool {
	return o.ID == 0
}

func (o *Topic) Equals(other *Topic) bool {
	return o.Name == other.Name
}

func (o *Topic) ToFormatInterface() format.Interface {
	return o.ToFormat()
}

func (o *Topic) ToFormat() *format.Topic {
	return &format.Topic{
		Common: format.NewCommon(o.ID),
		Name:   o.Name,
	}
}

func (o *Topic) FromFormat(topic *format.Topic) {
	*o = Topic{
		Topic: repo_model.Topic{
			ID:   topic.Index.GetID(),
			Name: topic.Name,
		},
	}
}

type TopicProvider struct {
	BaseProvider
}

func (o *TopicProvider) ToFormat(ctx context.Context, topic *Topic) *format.Topic {
	return topic.ToFormat()
}

func (o *TopicProvider) FromFormat(ctx context.Context, m *format.Topic) *Topic {
	var topic Topic
	topic.FromFormat(m)
	return &topic
}

func (o *TopicProvider) GetObjects(ctx context.Context, user *User, project *Project, page int) []*Topic {
	topics, _, err := repo_model.FindTopics(ctx, &repo_model.FindTopicOptions{
		ListOptions: db.ListOptions{Page: page, PageSize: o.g.perPage},
		RepoID:      project.GetID(),
	})
	if err != nil {
		panic(err)
	}

	return util.ConvertMap[*repo_model.Topic, *Topic](topics, TopicConverter)
}

func (o *TopicProvider) ProcessObject(ctx context.Context, user *User, project *Project, topic *Topic) {
}

func (o *TopicProvider) Get(ctx context.Context, user *User, project *Project, exemplar *Topic) *Topic {
	id := exemplar.GetID()
	topic, err := repo_model.GetRepoTopicByID(ctx, project.GetID(), id)
	if repo_model.IsErrTopicNotExist(err) {
		return &Topic{}
	}
	if err != nil {
		panic(err)
	}
	return TopicConverter(topic)
}

func (o *TopicProvider) Put(ctx context.Context, user *User, project *Project, topic, existing *Topic) *Topic {
	t, err := repo_model.AddTopic(ctx, project.GetID(), topic.Name)
	if err != nil {
		panic(err)
	}
	return o.Get(ctx, user, project, TopicConverter(t))
}

func (o *TopicProvider) Delete(ctx context.Context, user *User, project *Project, topic *Topic) *Topic {
	t := o.Get(ctx, user, project, topic)
	if !t.IsNil() {
		t, err := repo_model.DeleteTopic(ctx, project.GetID(), t.Name)
		if err != nil {
			panic(err)
		}
		return TopicConverter(t)
	}
	return t
}
