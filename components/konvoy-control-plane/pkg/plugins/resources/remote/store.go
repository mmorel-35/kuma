package remote

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/model"
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/model/rest"
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/store"
	util_http "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/util/http"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func NewStore(client util_http.Client, api rest.Api) store.ResourceStore {
	return &remoteStore{
		client: client,
		api:    api,
	}
}

var _ store.ResourceStore = &remoteStore{}

type remoteStore struct {
	client util_http.Client
	api    rest.Api
}

func (s *remoteStore) Create(ctx context.Context, res model.Resource, fs ...store.CreateOptionsFunc) error {
	opts := store.NewCreateOptions(fs...)
	meta := rest.ResourceMeta{
		Type: string(res.GetType()),
		Name: opts.Name,
		Mesh: opts.Mesh,
	}
	if err := s.upsert(ctx, res, meta); err != nil {
		return err
	}
	return nil
}
func (s *remoteStore) Update(ctx context.Context, res model.Resource, fs ...store.UpdateOptionsFunc) error {
	_ = store.NewUpdateOptions(fs...)
	meta := rest.ResourceMeta{
		Type: string(res.GetType()),
		Name: res.GetMeta().GetName(),
		Mesh: res.GetMeta().GetMesh(),
	}
	if err := s.upsert(ctx, res, meta); err != nil {
		return err
	}
	return nil
}

func (s *remoteStore) upsert(ctx context.Context, res model.Resource, meta rest.ResourceMeta) error {
	resourceApi, err := s.api.GetResourceApi(res.GetType())
	if err != nil {
		return errors.Wrapf(err, "failed to construct URI to fetch a list of %q", res.GetType())
	}
	restRes := rest.Resource{
		Meta: meta,
		Spec: res.GetSpec(),
	}
	b, err := json.Marshal(&restRes)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", resourceApi.Item(meta.Mesh, meta.Name), bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")
	statusCode, _, err := s.doRequest(ctx, req)
	if statusCode != http.StatusOK && statusCode != http.StatusCreated {
		return errors.Errorf("unexpected status code: %d", statusCode)
	}
	if err != nil {
		return err
	}
	res.SetMeta(rest.RemoteMeta{
		Namespace: "",
		Name:      meta.Name,
		Mesh:      meta.Mesh,
		Version:   "",
	})
	return nil
}
func (s *remoteStore) Delete(context.Context, model.Resource, ...store.DeleteOptionsFunc) error {
	return errors.Errorf("not implemented yet")
}
func (s *remoteStore) Get(ctx context.Context, res model.Resource, fs ...store.GetOptionsFunc) error {
	resourceApi, err := s.api.GetResourceApi(res.GetType())
	if err != nil {
		return errors.Wrapf(err, "failed to construct URI to fetch a list of %q", res.GetType())
	}
	opts := store.NewGetOptions(fs...)
	req, err := http.NewRequest("GET", resourceApi.Item(opts.Mesh, opts.Name), nil)
	if err != nil {
		return err
	}
	statusCode, b, err := s.doRequest(ctx, req)
	if err != nil {
		return err
	}
	if statusCode == 404 {
		return store.ErrorResourceNotFound(res.GetType(), opts.Namespace, opts.Name, opts.Mesh)
	}
	if statusCode != 200 {
		return errors.Errorf("unexpected status code: %d", statusCode)
	}
	return rest.Unmarshal(b, res)
}

func (s *remoteStore) List(ctx context.Context, rs model.ResourceList, fs ...store.ListOptionsFunc) error {
	resourceApi, err := s.api.GetResourceApi(rs.GetItemType())
	if err != nil {
		return errors.Wrapf(err, "failed to construct URI to fetch a list of %q", rs.GetItemType())
	}
	opts := store.NewListOptions(fs...)
	req, err := http.NewRequest("GET", resourceApi.List(opts.Mesh), nil)
	if err != nil {
		return err
	}
	statusCode, b, err := s.doRequest(ctx, req)
	if err != nil {
		return err
	}
	if statusCode != http.StatusOK {
		return errors.Errorf("unexpected status code: %d", statusCode)
	}
	return rest.UnmarshalList(b, rs)
}

// execute a request. Returns status code, body, error
func (s *remoteStore) doRequest(ctx context.Context, req *http.Request) (int, []byte, error) {
	req.Header.Set("Accept", "application/json")
	resp, err := s.client.Do(req.WithContext(ctx))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, b, nil
}
