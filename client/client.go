package client

import (
	"context"

	"github.com/razin99/cq-source-knowbe4/knowbe4client"
	"github.com/rs/zerolog"
)

type Client struct {
	logger  zerolog.Logger
	Spec    Spec
	KnowBe4 *knowbe4client.KnowBe4Client
}

func (c *Client) ID() string {
	return "knowbe4"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	c := Client{
		logger: logger,
		Spec:   *s,
		KnowBe4: &knowbe4client.KnowBe4Client{
			BaseURL: s.BaseURL,
			Token:   s.Token,
		},
	}
	return c, nil
}
