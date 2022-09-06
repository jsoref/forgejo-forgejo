// SPDX-License-Identifier: MIT

package driver

import (
	"context"

	"lab.forgefriends.org/friendlyforgeformat/gof3/forges/common"
	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
)

type BaseProvider struct {
	g *Forgejo
}

func (o *BaseProvider) SetForgejo(g *Forgejo) {
	o.g = g
}

type BaseProviderConstraint[Provider any] interface {
	*Provider
	SetForgejo(*Forgejo)
}

func NewProvider[T any, TPtr BaseProviderConstraint[T]](g *Forgejo) TPtr {
	var p TPtr
	p = new(T)
	p.SetForgejo(g)
	return p
}

func (o *BaseProvider) GetLocalMatchingRemote(ctx context.Context, format format.Interface, parents ...common.ContainerObjectInterface) (string, bool) {
	return "", false
}

type BaseProviderWithProjectProvider struct {
	BaseProvider
	project *ProjectProvider
}

func (o *BaseProviderWithProjectProvider) SetProjectProvider(project *ProjectProvider) {
	o.project = project
}

type BaseProviderWithProjectProviderConstraint[Provider any] interface {
	BaseProviderConstraint[Provider]
	SetProjectProvider(project *ProjectProvider)
}

func NewProviderWithProjectProvider[T any, TPtr BaseProviderWithProjectProviderConstraint[T]](g *Forgejo, project *ProjectProvider) TPtr {
	p := NewProvider[T, TPtr](g)
	p.SetProjectProvider(project)
	return p
}
