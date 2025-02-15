package compcontcron

import (
	"github.com/go-compcont/compcont-core"
)

const TypeID compcont.ComponentTypeID = "contrib.simple-cron-scheduler"

type Component interface {
	AddTask(taskName string, fn func())
}

type Config struct {
	Enabled bool                `ccf:"enabled"` // 是否启用调度
	Policy  map[string][]string `ccf:"policy"`  // 各个定时任务调度策略map[taskName]crontab
}

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, Component]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config Config) (instance Component, err error) {
		instance = New(config)
		return
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
