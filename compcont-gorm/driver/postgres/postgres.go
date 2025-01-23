package postgres

import (
	"github.com/go-compcont/compcont-core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const TypeID compcont.ComponentTypeID = "contrib.gorm.driver.postgres"

type Config struct {
	DSN                  string `ccf:"dsn"`
	WithoutQuotingCheck  bool   `ccf:"without_quoting_check"`
	PreferSimpleProtocol bool   `ccf:"prefer_simple_protocol"`
	WithoutReturning     bool   `ccf:"without_returning"`
}

func New(cc compcont.IComponentContainer, cfg Config) (c gorm.Dialector, err error) {
	c = postgres.New(postgres.Config{
		DSN:                  cfg.DSN,
		WithoutQuotingCheck:  cfg.WithoutQuotingCheck,
		PreferSimpleProtocol: cfg.PreferSimpleProtocol,
		WithoutReturning:     cfg.WithoutReturning,
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
