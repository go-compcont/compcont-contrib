package gorm

import (
	"github.com/go-compcont/compcont-core"
	"gorm.io/gorm"
)

const Type compcont.ComponentTypeID = "contrib.gorm"

type Config struct {
	Driver compcont.TypedComponentConfig[any, gorm.Dialector] `ccf:"driver"`
}

func Build(cc compcont.IComponentContainer, cfg Config) (c *gorm.DB, err error) {
	driverComp, err := cfg.Driver.LoadComponent(cc)
	if err != nil {
		return
	}
	db, err := gorm.Open(driverComp.Instance, &gorm.Config{})
	if err != nil {
		return
	}
	c = db
	return
}
