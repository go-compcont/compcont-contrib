package restyprovider

import (
	"context"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type Options struct {
	Ctx  context.Context
	Tags []string
	Url  *url.URL
}

type OptionsFunc func(o *Options)

func WithContext(ctx context.Context) OptionsFunc {
	return func(o *Options) {
		o.Ctx = ctx
	}
}

func WithURLOption(rawURL string) OptionsFunc {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return func(o *Options) {
		o.Url = u
	}
}

func WithTagOption(tag string) OptionsFunc {
	return func(o *Options) {
		o.Tags = append(o.Tags, tag)
	}
}

type RestyProvider interface {
	GetResty(opts ...OptionsFunc) (*resty.Client, error)
}

type GetRestyFunc func(opts ...OptionsFunc) (*resty.Client, error)

func (m GetRestyFunc) GetResty(opts ...OptionsFunc) (*resty.Client, error) {
	return m(opts...)
}
