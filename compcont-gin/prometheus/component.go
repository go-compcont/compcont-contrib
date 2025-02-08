package prom

import (
	"github.com/gin-gonic/gin"
	"github.com/go-compcont/compcont-core"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const TypeID = "contrib.gin.prometheus"

type Config struct {
	Gin compcont.TypedComponentConfig[any, gin.IRouter] `ccf:"gin"`
}

func New(cc compcont.IComponentContainer, cfg Config) (reg *prometheus.Registry, err error) {
	routerComp := cfg.Gin.MustLoadComponent(cc)
	router := routerComp.Instance

	reg = prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewBuildInfoCollector(),
		collectors.NewGoCollector(
			collectors.WithGoCollectorRuntimeMetrics(collectors.MetricsAll),
		),
	)
	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	})

	router.GET("/metrics", gin.WrapH(handler))
	return
}

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, prometheus.Registerer]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config Config) (instance prometheus.Registerer, err error) {
		return New(ctx.Container, config)
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
