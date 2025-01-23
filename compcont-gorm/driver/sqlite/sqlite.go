package sqlite

import (
	"github.com/go-compcont/compcont-core"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const TypeID compcont.ComponentTypeID = "contrib.gorm.driver.sqlite"

type Config struct {
	DSN string `ccf:"dsn"`
}

func New(cc compcont.IComponentContainer, cfg Config) (c gorm.Dialector, err error) {
	c = sqlite.New(sqlite.Config{
		DSN: cfg.DSN,
	})
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
