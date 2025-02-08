package gorm

import (
	"github.com/go-compcont/compcont-core"
	"gorm.io/gorm"
)

const TypeID compcont.ComponentTypeID = "contrib.gorm"

type Config struct {
	Driver compcont.TypedComponentConfig[any, gorm.Dialector] `ccf:"driver"`
}

func New(cc compcont.IComponentContainer, cfg Config) (c *gorm.DB, err error) {
	driverComp := cfg.Driver.MustLoadComponent(cc)
	db, err := gorm.Open(driverComp.Instance, &gorm.Config{})
	if err != nil {
		return
	}
	c = db
	return
}

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, any]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config Config) (instance any, err error) {
		return New(ctx.Container, config)
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
