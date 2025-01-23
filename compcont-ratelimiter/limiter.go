package compcontratelimiter

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter interface {
	Wait(ctx context.Context, key string)
	Reserve(key string)
}

type rateLimiterImpl struct {
	Config
	m  map[string]*rate.Limiter
	mu sync.RWMutex
}

func newImpl(cfg Config) RateLimiter {
	return &rateLimiterImpl{
		Config: cfg,
		m:      make(map[string]*rate.Limiter),
	}
}

func (p *rateLimiterImpl) getRateLimiter(key string) *rate.Limiter {
	p.mu.RLock()
	if v, ok := p.m[key]; ok {
		p.mu.RUnlock()
		return v
	}
	p.mu.RUnlock()

	limiterConfig := p.DefaultLimiter
	if v, ok := p.SpecialLimiter[key]; ok {
		limiterConfig = v
	}
	limiter := rate.NewLimiter(rate.Every(limiterConfig.TokenInterval), limiterConfig.Bursts)
	p.mu.Lock()
	p.m[key] = limiter
	p.mu.Unlock()
	return limiter
}

func (p *rateLimiterImpl) Reserve(key string) {
	p.getRateLimiter(key).Reserve()
}

func (p *rateLimiterImpl) Wait(ctx context.Context, key string) {
	p.getRateLimiter(key).Wait(ctx)
}

type LimiterConfig struct {
	Bursts        int           `ccf:"bursts"`
	TokenInterval time.Duration `ccf:"token_interval"`
}

type Config struct {
	DefaultLimiter LimiterConfig            `ccf:"default_limiter"` // 默认限流器
	SpecialLimiter map[string]LimiterConfig `ccf:"special_limiter"` // 特殊key的限流器
}
