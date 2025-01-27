package compcontratelimiter

import (
	"context"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"golang.org/x/time/rate"
)

// RateLimiter 定义速率限制器接口
type RateLimiter interface {
	Wait(ctx context.Context, key string)
	Reserve(key string)
}

// rateLimiterImpl 实现 RateLimiter 接口
type rateLimiterImpl struct {
	Config
	cache *lru.Cache // 使用 LRU 缓存存储速率限制器
}

// newImpl 初始化一个新的速率限制器
func newImpl(cfg Config) (RateLimiter, error) {
	// 如果未配置 LRUSize，则使用默认值 1024
	if cfg.LRUSize <= 0 {
		cfg.LRUSize = 1024
	}

	// 创建一个指定大小的 LRU 缓存
	cache, err := lru.New(cfg.LRUSize)
	if err != nil {
		return nil, err
	}

	return &rateLimiterImpl{
		Config: cfg,
		cache:  cache,
	}, nil
}

// getRateLimiter 获取或创建速率限制器
func (p *rateLimiterImpl) getRateLimiter(key string) *rate.Limiter {
	// 从缓存中获取速率限制器
	if v, ok := p.cache.Get(key); ok {
		return v.(*rate.Limiter)
	}

	// 如果缓存中不存在，则根据配置创建新的速率限制器
	limiterConfig := p.DefaultLimiter
	if v, ok := p.SpecialLimiter[key]; ok {
		limiterConfig = v
	}
	limiter := rate.NewLimiter(rate.Every(limiterConfig.TokenInterval), limiterConfig.Bursts)

	// 将新的速率限制器添加到缓存中
	p.cache.Add(key, limiter)
	return limiter
}

// Reserve 为指定键预留一个令牌
func (p *rateLimiterImpl) Reserve(key string) {
	p.getRateLimiter(key).Reserve()
}

// Wait 阻塞等待，直到满足指定键的速率限制条件
func (p *rateLimiterImpl) Wait(ctx context.Context, key string) {
	p.getRateLimiter(key).Wait(ctx)
}

// LimiterConfig 定义速率限制器的配置
type LimiterConfig struct {
	Bursts        int           `ccf:"bursts"`         // 令牌桶容量
	TokenInterval time.Duration `ccf:"token_interval"` // 令牌生成间隔
}

// Config 定义速率限制器的全局配置
type Config struct {
	DefaultLimiter LimiterConfig            `ccf:"default_limiter"` // 默认速率限制器配置
	SpecialLimiter map[string]LimiterConfig `ccf:"special_limiter"` // 特殊键的速率限制器配置
	LRUSize        int                      `ccf:"lru_size"`        // LRU 缓存大小
}
