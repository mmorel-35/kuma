package tokens

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	error_types "github.com/kumahq/kuma/v2/pkg/core/rest/errors/types"
	util_http "github.com/kumahq/kuma/v2/pkg/util/http"
)

type TokenClient struct {
	client util_http.Client
	url    string
}

func NewTokenClient(client util_http.Client, entity string) TokenClient {
	return TokenClient{
		client: client,
		url:    "/tokens/" + entity,
	}
}

func (tc TokenClient) Generate(tokenReq any) (string, error) {
	reqBytes, err := json.Marshal(tokenReq)
	if err != nil {
		return "", fmt.Errorf("could not marshal token request to json: %w", err)
	}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, tc.url, bytes.NewReader(reqBytes))
	if err != nil {
		return "", fmt.Errorf("could not construct the request: %w", err)
	}
	req.Header.Set("content-type", "application/json")
	resp, err := tc.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not execute the request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read a body of the request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		var kumaErr error_types.Error
		if err := json.Unmarshal(body, &kumaErr); err == nil {
			if kumaErr.Title != "" {
				return "", &kumaErr
			}
		}
		return "", fmt.Errorf("(%d): %s", resp.StatusCode, body)
	}
	return string(body), nil
}
