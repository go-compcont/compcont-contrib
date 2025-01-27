package compcontredis

import (
	"context"
	"io"

	"github.com/go-compcont/compcont-core"
	"github.com/redis/go-redis/v9"
)

const TypeID compcont.ComponentTypeID = "contrib.redis"

type Config struct {
	URL string `ccf:"url"`
}

type Component interface {
	redis.Cmdable
	io.Closer
}

func New(cfg Config) (comp Component, err error) {
	options, err := redis.ParseURL(cfg.URL)
	if err != nil {
		return
	}
	comp = redis.NewClient(options)
	status := comp.Ping(context.Background())
	if err = status.Err(); err != nil {
		return
	}
	return
}

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, Component]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config Config) (instance Component, err error) {
		return New(config)
	},
	DestroyInstanceFunc: func(ctx compcont.BuildContext, instance Component) (err error) {
		return instance.Close()
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
