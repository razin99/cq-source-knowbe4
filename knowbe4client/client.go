package knowbe4client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type KnowBe4Client struct {
	BaseURL    string
	Token      string
	HttpClient http.Client
}

func (kb4 KnowBe4Client) NewRequest(ctx context.Context, path string, params url.Values) (*http.Response, error) {
	u, err := url.Parse(kb4.BaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = path
	u.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", kb4.Token))
	req.Header.Set("Accept", "application/json")

	res, err := kb4.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Error: %d, Path: %s", res.StatusCode, req.URL.String())
	}

	return res, nil
}
