package gorm

import (
	"context"
	"time"

	"github.com/go-compcont/compcont-core"
	compcontzap "github.com/go-compcont/compcont-std/compcont-zap"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const TypeID compcont.ComponentTypeID = "contrib.gorm"

type Config struct {
	Driver compcont.TypedComponentConfig[any, gorm.Dialector] `ccf:"driver"`
}

type log struct{}

func (l *log) LogMode(lv logger.LogLevel) logger.Interface {
	return l
}

func (l *log) Info(ctx context.Context, format string, args ...interface{}) {
	compcontzap.FromContext(ctx).Sugar().Infof(format, args...)
}

func (l *log) Warn(ctx context.Context, format string, args ...interface{}) {
	compcontzap.FromContext(ctx).Sugar().Warnf(format, args...)
}

func (l *log) Error(ctx context.Context, format string, args ...interface{}) {
	compcontzap.FromContext(ctx).Sugar().Errorf(format, args...)
}

func (l *log) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	compcontzap.FromContext(ctx).Debug("Gorm Trace",
		zap.Time("begin", begin),
		zap.String("sql", sql),
		zap.Int64("rowsAffected", rowsAffected),
		zap.Error(err),
	)
}

func New(cc compcont.IComponentContainer, cfg Config) (c *gorm.DB, err error) {
	driverComp := cfg.Driver.MustLoadComponent(cc)
	db, err := gorm.Open(driverComp.Instance, &gorm.Config{
		Logger: &log{},
	})
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
