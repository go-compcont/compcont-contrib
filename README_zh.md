# compcont-contrib

[English](README.md) | [中文](README_zh.md)

这是 [compcont-core](https://github.com/go-compcont/compcont-core) 组件容器框架的扩展组件集合。

## 概述

`compcont-contrib` 提供了一系列开箱即用的组件，将常用的 Go 库集成到 compcont 框架中。这些组件可以通过声明式配置来使用，并可以组合在一起构建应用程序。

## 安装

```bash
go get github.com/go-compcont/compcont-contrib
```

预加载所有组件，导入根包：

```go
import _ "github.com/go-compcont/compcont-contrib"
```

或者只导入需要的特定组件：

```go
import _ "github.com/go-compcont/compcont-contrib/compcont-gin/gin"
import _ "github.com/go-compcont/compcont-contrib/compcont-redis"
```

## 可用组件

### Web 框架

| 组件 | 类型 ID | 描述 |
|------|---------|------|
| [compcont-gin/gin](compcont-gin/gin) | `contrib.gin` | Gin HTTP Web 框架 |
| [compcont-gin/pprof](compcont-gin/pprof) | `contrib.gin.pprof` | Gin 的 pprof 性能分析端点 |
| [compcont-gin/prometheus](compcont-gin/prometheus) | `contrib.gin.prometheus` | Gin 的 Prometheus 指标端点 |

### Gin 中间件

| 组件 | 类型 ID | 描述 |
|------|---------|------|
| [compcont-gin/middleware/cors](compcont-gin/middleware/cors) | `contrib.gin.middleware.cors` | CORS 跨域中间件 |
| [compcont-gin/middleware/prometheus](compcont-gin/middleware/prometheus) | `contrib.gin.middleware.prometheus` | Prometheus 指标收集中间件 |
| [compcont-gin/middleware/recovery](compcont-gin/middleware/recovery) | `contrib.gin.middleware.recovery` | Panic 恢复中间件 |
| [compcont-gin/middleware/zap](compcont-gin/middleware/zap) | `contrib.gin.middleware.zap` | Zap 日志中间件，支持请求/响应记录 |

### 数据库

| 组件 | 类型 ID | 描述 |
|------|---------|------|
| [compcont-gorm/gorm](compcont-gorm/gorm) | `contrib.gorm` | GORM ORM，支持上下文感知日志 |
| [compcont-gorm/driver/postgres](compcont-gorm/driver/postgres) | `contrib.gorm.driver.postgres` | GORM 的 PostgreSQL 驱动 |
| [compcont-gorm/driver/sqlite](compcont-gorm/driver/sqlite) | `contrib.gorm.driver.sqlite` | GORM 的 SQLite 驱动 |
| [compcont-redis](compcont-redis) | `contrib.redis` | Redis 客户端 |

### 云服务

| 组件 | 类型 ID | 描述 |
|------|---------|------|
| [compcont-s3](compcont-s3) | `contrib.s3` | AWS S3 客户端，支持灵活配置 |

### HTTP 客户端

| 组件 | 类型 ID | 描述 |
|------|---------|------|
| [compcont-resty](compcont-resty) | `contrib.resty-provider-simple` | 简单的 resty HTTP 客户端，支持重试、代理和 TLS |
| [compcont-resty](compcont-resty) | `contrib.resty-provider-rule` | 基于规则的 resty HTTP 客户端，支持表达式匹配 |

### 工具类

| 组件 | 类型 ID | 描述 |
|------|---------|------|
| [compcont-cron](compcont-cron) | `contrib.simple-cron-scheduler` | 定时任务调度器，支持手动触发 |
| [compcont-jwt](compcont-jwt) | `contrib.jwt` | JWT 认证工具 |
| [compcont-ratelimiter](compcont-ratelimiter) | `contrib.ratelimiter` | 令牌桶限流器，带 LRU 缓存 |
| [compcont-ddddocr](compcont-ddddocr) | `contrib.ddddocr` | ddddocr 服务的 OCR 客户端 |
| [compcont-graph](compcont-graph) | `contrib.compcont-graph` | 组件关系图可视化（PlantUML 导出） |

## 快速开始

### 示例：设置带中间件的 Gin HTTP 服务器

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

### 示例：Redis 连接

```yaml
components:
  - name: redis
    type: contrib.redis
    config:
      url: "redis://localhost:6379/0"
```

### 示例：GORM 与 PostgreSQL

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

### 示例：JWT 认证

```yaml
components:
  - name: jwt_auth
    type: contrib.jwt
    config:
      secret_key: "your-secret-key"
```

### 示例：限流器

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

### 示例：定时任务调度器

```yaml
components:
  - name: scheduler
    type: contrib.simple-cron-scheduler
    config:
      enabled: true
      policy:
        cleanup_task:
          - "0 0 * * *"  # 每天午夜执行
          - "manual"     # 同时支持手动触发
```

## 组件接口参考

### Gin 组件
```go
type Component interface {
    gin.IRouter
}
```

### Redis 组件
```go
type Component interface {
    redis.Cmdable
    io.Closer
}
```

### JWT 组件
```go
type JWTAuther interface {
    Verify(token string) bool
    Parse(token string, payload any) error
    Generate(payload any) (string, error)
}
```

### 限流器组件
```go
type RateLimiter interface {
    Wait(ctx context.Context, key string)
    Reserve(key string)
}
```

### 定时任务组件
```go
type Component interface {
    AddTask(taskName string, fn func(ctx context.Context) error)
    DoTask(ctx context.Context, taskName string) error
}
```

### Resty 提供者组件
```go
type RestyProvider interface {
    GetResty(opts ...OptionsFunc) (*resty.Client, error)
}
```

## 许可证

本项目是开源的。有关具体许可信息，请参阅各个组件包。

## 贡献

欢迎贡献！请随时提交问题和拉取请求。
