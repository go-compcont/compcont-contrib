# compcont-contrib

[English](README.md) | [中文](README_zh.md)

A collection of contributed components for the [compcont-core](https://github.com/go-compcont/compcont-core) component container framework.

## Overview

`compcont-contrib` provides a rich set of ready-to-use components that integrate popular Go libraries into the compcont framework. These components can be configured declaratively and composed together to build applications.

## Installation

```bash
go get github.com/go-compcont/compcont-contrib
```

To preload all components, import the root package:

```go
import _ "github.com/go-compcont/compcont-contrib"
```

Or import only the specific components you need:

```go
import _ "github.com/go-compcont/compcont-contrib/compcont-gin/gin"
import _ "github.com/go-compcont/compcont-contrib/compcont-redis"
```

## Available Components

### Web Framework

| Component | Type ID | Description |
|-----------|---------|-------------|
| [compcont-gin/gin](compcont-gin/gin) | `contrib.gin` | Gin HTTP web framework |
| [compcont-gin/pprof](compcont-gin/pprof) | `contrib.gin.pprof` | pprof profiling endpoints for Gin |
| [compcont-gin/prometheus](compcont-gin/prometheus) | `contrib.gin.prometheus` | Prometheus metrics endpoint for Gin |

### Gin Middlewares

| Component | Type ID | Description |
|-----------|---------|-------------|
| [compcont-gin/middleware/cors](compcont-gin/middleware/cors) | `contrib.gin.middleware.cors` | CORS middleware |
| [compcont-gin/middleware/prometheus](compcont-gin/middleware/prometheus) | `contrib.gin.middleware.prometheus` | Prometheus metrics middleware |
| [compcont-gin/middleware/recovery](compcont-gin/middleware/recovery) | `contrib.gin.middleware.recovery` | Panic recovery middleware |
| [compcont-gin/middleware/zap](compcont-gin/middleware/zap) | `contrib.gin.middleware.zap` | Zap logging middleware with request/response recording |

### Database

| Component | Type ID | Description |
|-----------|---------|-------------|
| [compcont-gorm/gorm](compcont-gorm/gorm) | `contrib.gorm` | GORM ORM with context-aware logging |
| [compcont-gorm/driver/postgres](compcont-gorm/driver/postgres) | `contrib.gorm.driver.postgres` | PostgreSQL driver for GORM |
| [compcont-gorm/driver/sqlite](compcont-gorm/driver/sqlite) | `contrib.gorm.driver.sqlite` | SQLite driver for GORM |
| [compcont-redis](compcont-redis) | `contrib.redis` | Redis client |

### Cloud Services

| Component | Type ID | Description |
|-----------|---------|-------------|
| [compcont-s3](compcont-s3) | `contrib.s3` | AWS S3 client with flexible configuration |

### HTTP Client

| Component | Type ID | Description |
|-----------|---------|-------------|
| [compcont-resty](compcont-resty) | `contrib.resty-provider-simple` | Simple resty HTTP client with retry, proxy, and TLS support |
| [compcont-resty](compcont-resty) | `contrib.resty-provider-rule` | Rule-based resty HTTP client with expression matching |

### Utilities

| Component | Type ID | Description |
|-----------|---------|-------------|
| [compcont-cron](compcont-cron) | `contrib.simple-cron-scheduler` | Cron job scheduler with manual trigger support |
| [compcont-jwt](compcont-jwt) | `contrib.jwt` | JWT authentication utilities |
| [compcont-ratelimiter](compcont-ratelimiter) | `contrib.ratelimiter` | Token bucket rate limiter with LRU cache |
| [compcont-ddddocr](compcont-ddddocr) | `contrib.ddddocr` | OCR client for ddddocr service |
| [compcont-graph](compcont-graph) | `contrib.compcont-graph` | Component graph visualization (PlantUML export) |

## Quick Start

### Example: Setting up a Gin HTTP Server with Middlewares

```yaml
components:
  - name: gin
    type: contrib.gin
    config:
      mode: release
      listen_addrs:
        - ":8080"
      middlewares:
        - type: contrib.gin.middleware.recovery
        - type: contrib.gin.middleware.cors
          config:
            allow_all_origins: true

  - name: pprof
    type: contrib.gin.pprof
    deps: [gin]
    config:
      gin:
        refer: ../gin
      route_prefix: /debug/pprof
```

### Example: Redis Connection

```yaml
components:
  - name: redis
    type: contrib.redis
    config:
      url: "redis://localhost:6379/0"
```

### Example: GORM with PostgreSQL

```yaml
components:
  - name: postgres_driver
    type: contrib.gorm.driver.postgres
    config:
      dsn: "host=localhost user=postgres password=secret dbname=myapp port=5432"

  - name: db
    type: contrib.gorm
    deps: [postgres_driver]
    config:
      driver:
        refer: ../postgres_driver
```

### Example: JWT Authentication

```yaml
components:
  - name: jwt_auth
    type: contrib.jwt
    config:
      secret_key: "your-secret-key"
```

### Example: Rate Limiter

```yaml
components:
  - name: rate_limiter
    type: contrib.ratelimiter
    config:
      default_limiter:
        bursts: 10
        token_interval: 100ms
      special_limiter:
        api_key_1:
          bursts: 100
          token_interval: 10ms
      lru_size: 1024
```

### Example: Cron Scheduler

```yaml
components:
  - name: scheduler
    type: contrib.simple-cron-scheduler
    config:
      enabled: true
      policy:
        cleanup_task:
          - "0 0 * * *"  # Daily at midnight
          - "manual"     # Also allow manual trigger
```

## Component Interface Reference

### Gin Component
```go
type Component interface {
    gin.IRouter
}
```

### Redis Component
```go
type Component interface {
    redis.Cmdable
    io.Closer
}
```

### JWT Component
```go
type JWTAuther interface {
    Verify(token string) bool
    Parse(token string, payload any) error
    Generate(payload any) (string, error)
}
```

### Rate Limiter Component
```go
type RateLimiter interface {
    Wait(ctx context.Context, key string)
    Reserve(key string)
}
```

### Cron Component
```go
type Component interface {
    AddTask(taskName string, fn func(ctx context.Context) error)
    DoTask(ctx context.Context, taskName string) error
}
```

### Resty Provider Component
```go
type RestyProvider interface {
    GetResty(opts ...OptionsFunc) (*resty.Client, error)
}
```

## License

This project is open source. See the individual component packages for specific licensing information.

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.
