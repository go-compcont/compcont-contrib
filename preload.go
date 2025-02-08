package compcontcontrib

import (
	_ "github.com/go-compcont/compcont-contrib/compcont-cron"
	_ "github.com/go-compcont/compcont-contrib/compcont-ddddocr"
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/gin"
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/middleware/cors"
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/middleware/prometheus"
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/middleware/recovery"
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/middleware/zap"
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/pprof"
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/prometheus"
	_ "github.com/go-compcont/compcont-contrib/compcont-gorm"
	_ "github.com/go-compcont/compcont-contrib/compcont-graph"
	_ "github.com/go-compcont/compcont-contrib/compcont-jwt"
	_ "github.com/go-compcont/compcont-contrib/compcont-ratelimiter"
	_ "github.com/go-compcont/compcont-contrib/compcont-redis"
	_ "github.com/go-compcont/compcont-contrib/compcont-resty"
	_ "github.com/go-compcont/compcont-contrib/compcont-s3"
)
