package compcontratelimiter

import "github.com/go-compcont/compcont-core"

const TypeID compcont.ComponentTypeID = "contrib.ratelimiter"

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, RateLimiter]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config Config) (instance RateLimiter, err error) {
		instance, err = newImpl(config)
		return
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
