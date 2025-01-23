package sqlite

import (
	"github.com/go-compcont/compcont-core"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const Type compcont.ComponentTypeID = "contrib.gorm.driver.postgres"

type Config struct {
	DSN string `ccf:"dsn"`
}

func Build(cc compcont.IComponentContainer, cfg Config) (c gorm.Dialector, err error) {
	sqlite.New(sqlite.Config{
		DSN: cfg.DSN,
	})

	c = sqlite.New(sqlite.Config{
		DSN: cfg.DSN,
	})
	return
}
