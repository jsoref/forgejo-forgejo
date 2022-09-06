// SPDX-License-Identifier: MIT

package driver

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"code.gitea.io/gitea/models/db"
	repo_model "code.gitea.io/gitea/models/repo"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/services/attachment"

	"github.com/google/uuid"
	"lab.forgefriends.org/friendlyforgeformat/gof3/format"
	"lab.forgefriends.org/friendlyforgeformat/gof3/util"
)

type Asset struct {
	repo_model.Attachment
	DownloadFunc func() io.ReadCloser
}

func AssetConverter(f *repo_model.Attachment) *Asset {
	return &Asset{
		Attachment: *f,
	}
}

func (o Asset) GetID() int64 {
	return o.ID
}

func (o Asset) GetIDString() string {
	return fmt.Sprintf("%d", o.GetID())
}

func (o *Asset) SetID(id int64) {
	o.ID = id
}

func (o *Asset) SetIDString(id string) {
	o.SetID(util.ParseInt(id))
}

func (o *Asset) IsNil() bool {
	return o.ID == 0
}

func (o *Asset) Equals(other *Asset) bool {
	return o.Name == other.Name
}

func (o *Asset) ToFormatInterface() format.Interface {
	return o.ToFormat()
}

func (o *Asset) ToFormat() *format.ReleaseAsset {
	return &format.ReleaseAsset{
		Common:        format.NewCommon(o.ID),
		Name:          o.Name,
		Size:          int(o.Size),
		DownloadCount: int(o.DownloadCount),
		Created:       o.CreatedUnix.AsTime(),
		DownloadURL:   o.DownloadURL(),
		DownloadFunc:  o.DownloadFunc,
	}
}

func (o *Asset) FromFormat(asset *format.ReleaseAsset) {
	*o = Asset{
		Attachment: repo_model.Attachment{
			ID:                asset.GetID(),
			Name:              asset.Name,
			Size:              int64(asset.Size),
			DownloadCount:     int64(asset.DownloadCount),
			CustomDownloadURL: asset.DownloadURL,
			CreatedUnix:       timeutil.TimeStamp(asset.Created.Unix()),
		},
		DownloadFunc: asset.DownloadFunc,
	}
}

type AssetProvider struct {
	BaseProvider
}

func (o *AssetProvider) ToFormat(ctx context.Context, asset *Asset) *format.ReleaseAsset {
	httpClient := o.g.GetNewMigrationHTTPClient()()
	a := asset.ToFormat()
	a.DownloadFunc = func() io.ReadCloser {
		o.g.GetLogger().Debug("download from %s", asset.DownloadURL())
		req, err := http.NewRequest("GET", asset.DownloadURL(), nil)
		if err != nil {
			panic(err)
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			panic(fmt.Errorf("while downloading %s %w", asset.DownloadURL(), err))
		}

		// resp.Body is closed by the consumer
		return resp.Body
	}
	return a
}

func (o *AssetProvider) FromFormat(ctx context.Context, p *format.ReleaseAsset) *Asset {
	var asset Asset
	asset.FromFormat(p)
	return &asset
}

func (o *AssetProvider) ProcessObject(ctx context.Context, user *User, project *Project, release *Release, asset *Asset) {
}

func (o *AssetProvider) GetObjects(ctx context.Context, user *User, project *Project, release *Release, page int) []*Asset {
	if page > 1 {
		return []*Asset{}
	}
	r, err := repo_model.GetReleaseByID(ctx, release.GetID())
	if err != nil {
		panic(err)
	}
	if err := r.LoadAttributes(ctx); err != nil {
		panic(fmt.Errorf("error while listing assets: %v", err))
	}

	return util.ConvertMap[*repo_model.Attachment, *Asset](r.Attachments, AssetConverter)
}

func (o *AssetProvider) Get(ctx context.Context, user *User, project *Project, release *Release, exemplar *Asset) *Asset {
	id := exemplar.GetID()
	asset, err := repo_model.GetAttachmentByID(ctx, id)
	if repo_model.IsErrAttachmentNotExist(err) {
		return &Asset{}
	}
	if err != nil {
		panic(err)
	}
	return AssetConverter(asset)
}

func (o *AssetProvider) Put(ctx context.Context, user *User, project *Project, release *Release, asset, existing *Asset) *Asset {
	var result *Asset

	if existing == nil || existing.IsNil() {
		a := asset.Attachment
		a.UploaderID = user.GetID()
		a.RepoID = project.GetID()
		a.ReleaseID = release.GetID()
		a.UUID = uuid.New().String()

		download := asset.DownloadFunc()
		defer download.Close()

		insertedAttachment, err := attachment.NewAttachment(ctx, &a, download, asset.Size)
		if err != nil {
			panic(err)
		}
		result = AssetConverter(insertedAttachment)
	} else {
		var u repo_model.Attachment
		u.ID = existing.GetID()
		cols := make([]string, 0, 10)

		if asset.Name != existing.Name {
			u.Name = asset.Name
			cols = append(cols, "name")
		}
		if len(cols) > 0 {
			if _, err := db.GetEngine(ctx).ID(existing.ID).Cols(cols...).Update(u); err != nil {
				panic(err)
			}
		}
		result = existing
	}
	return o.Get(ctx, user, project, release, result)
}

func (o *AssetProvider) Delete(ctx context.Context, user *User, project *Project, release *Release, asset *Asset) *Asset {
	a := o.Get(ctx, user, project, release, asset)
	if !a.IsNil() {
		err := repo_model.DeleteAttachment(ctx, &a.Attachment, true)
		if err != nil {
			panic(err)
		}
	}
	return a
}
