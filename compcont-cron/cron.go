package compcontcron

import (
	"context"
	"errors"
	"time"

	compcontzap "github.com/go-compcont/compcont-std/compcont-zap"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type componentImpl struct {
	cron   *cron.Cron
	policy map[string][]string
	fnMap  map[string]func(ctx context.Context) error
	logger *zap.Logger
}

func New(logger *zap.Logger, cfg Config) *componentImpl {
	c := cron.New()
	if cfg.Enabled {
		c.Start()
	}
	return &componentImpl{
		logger: logger,
		cron:   c,
		policy: cfg.Policy,
		fnMap:  make(map[string]func(ctx context.Context) error),
	}
}

func (c *componentImpl) DoTask(ctx context.Context, taskName string) (err error) {
	if fn, ok := c.fnMap[taskName]; !ok {
		return errors.New("task not found: " + taskName)
	} else {
		logger := c.logger.With(
			zap.String("taskName", taskName),
			zap.String("taskStartTime", time.Now().Format(time.RFC3339)),
			zap.String("taskPolicy", "manual"),
		)

		logger.Info("do task")
		if err = fn(ctx); err != nil {
			logger.Error("task run failed", zap.Error(err))
			return
		}
		logger.Info("do task finished", zap.String("taskEndTime", time.Now().Format(time.RFC3339)))
	}
	return
}

func (c *componentImpl) AddTask(taskName string, fn func(ctx context.Context) error) {
	c.logger.Info("AddTask", zap.String("taskName", taskName))
	c.fnMap[taskName] = fn
	if taskPolicy, ok := c.policy[taskName]; ok {
		for _, policy := range taskPolicy {
			if policy == "manual" {
				continue
			}
			c.cron.AddFunc(policy, func() {
				logger := c.logger.With(
					zap.String("taskName", taskName),
					zap.String("taskStartTime", time.Now().Format(time.RFC3339)),
					zap.String("taskPolicy", policy),
				)
				err := fn(compcontzap.WithContext(context.Background(), logger))
				if err != nil {
					logger.Error("task run failed", zap.Error(err))
					return
				}
				logger.Info("do task finished", zap.String("taskEndTime", time.Now().Format(time.RFC3339)))
			})
		}
		return
	}
}
