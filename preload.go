package compcontcontrib

import (
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/gin"
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/middleware/prometheus"
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/middleware/recovery"
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/middleware/zap"
	_ "github.com/go-compcont/compcont-contrib/compcont-gin/pprof"
	_ "github.com/go-compcont/compcont-contrib/compcont-jwt"
	_ "github.com/go-compcont/compcont-contrib/compcont-redis"
	_ "github.com/go-compcont/compcont-contrib/compcont-resty"
	_ "github.com/go-compcont/compcont-contrib/compcont-s3"
)
