package compcontcron

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-compcont/compcont-core"
	"go.uber.org/zap"
)

const TypeID compcont.ComponentTypeID = "contrib.simple-cron-scheduler"

type Component interface {
	AddTask(taskName string, fn func(ctx context.Context) error)
	DoTask(ctx context.Context, taskName string) error
}

type Config struct {
	Logger      *compcont.TypedComponentConfig[any, *zap.Logger] `ccf:"logger"`        // 任务日志
	Enabled     bool                                             `ccf:"enabled"`       // 是否启用自动调度
	Policy      map[string][]string                              `ccf:"policy"`        // 各个定时任务调度策略map[taskName]crontab
	Gin         *compcont.TypedComponentConfig[any, gin.IRouter] `ccf:"gin"`           // 管理接口
	GinDoRouter string                                           `ccf:"gin_do_router"` // 任务执行接口
}

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, Component]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config Config) (instance Component, err error) {
		var logger *zap.Logger
		if config.Logger == nil {
			logger, err = zap.NewDevelopment()
			if err != nil {
				return
			}
		} else {
			logger = config.Logger.MustLoadComponent(ctx.Container).Instance
		}
		instance = New(logger, config)
		if config.Gin != nil {
			ginComp := config.Gin.MustLoadComponent(ctx.Container).Instance
			ginComp.GET(config.GinDoRouter, func(ctx *gin.Context) {
				err := instance.DoTask(ctx, ctx.Query("task_name"))
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					return
				}
			})
		}
		return
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
