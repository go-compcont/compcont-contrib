package postgres

import (
	"github.com/go-compcont/compcont-core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const Type compcont.ComponentTypeID = "contrib.gorm.driver.postgres"

type Config struct {
	DSN                  string `ccf:"dsn"`
	WithoutQuotingCheck  bool   `ccf:"without_quoting_check"`
	PreferSimpleProtocol bool   `ccf:"prefer_simple_protocol"`
	WithoutReturning     bool   `ccf:"without_returning"`
}

func Build(cc compcont.IComponentContainer, cfg Config) (c gorm.Dialector, err error) {
	c = postgres.New(postgres.Config{
		DSN:                  cfg.DSN,
		WithoutQuotingCheck:  cfg.WithoutQuotingCheck,
		PreferSimpleProtocol: cfg.PreferSimpleProtocol,
		WithoutReturning:     cfg.WithoutReturning,
	})
	return
}
