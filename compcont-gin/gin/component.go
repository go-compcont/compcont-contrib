package compcontgin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-compcont/compcont-core"
)

type Config struct {
	Mode        string                                                `ccf:"mode"`
	ListenAddrs []string                                              `ccf:"listen_addrs"`
	Middlewares []compcont.TypedComponentConfig[any, gin.HandlerFunc] `ccf:"middlewares"`
}

type Component interface {
	gin.IRouter
}

func New(cc compcont.IComponentContainer, cfg Config) (c Component, err error) {
	gin.SetMode(cfg.Mode)
	g := gin.New(func(e *gin.Engine) { e.ContextWithFallback = true })
	var middlewares []gin.HandlerFunc
	for _, middlewareCfg := range cfg.Middlewares {
		middlewares = append(middlewares, middlewareCfg.MustLoadComponent(cc).Instance)
	}
	g.Use(middlewares...)
	go g.Run(cfg.ListenAddrs...)
	c = g
	return
}

const TypeID compcont.ComponentTypeID = "contrib.gin"

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, Component]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config Config) (instance Component, err error) {
		return New(ctx.Container, config)
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
