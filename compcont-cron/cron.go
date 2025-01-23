package compcontcron

import "github.com/robfig/cron/v3"

type componentImpl struct {
	cron   *cron.Cron
	policy map[string]string
}

func New(cfg Config) *componentImpl {
	c := cron.New()
	if cfg.Enabled {
		c.Start()
	}
	return &componentImpl{cron: c, policy: cfg.Policy}
}

func (c *componentImpl) AddTask(taskName string, fn func()) {
	if taskPolicy, ok := c.policy[taskName]; ok {
		c.cron.AddFunc(taskPolicy, fn)
	}
}
