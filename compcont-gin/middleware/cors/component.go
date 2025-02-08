package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-compcont/compcont-core"
)

const TypeID compcont.ComponentTypeID = "contrib.gin.middleware.recovery"

type Config struct {
	AllowAllOrigins           bool          `ccf:"allow_all_origins"`
	AllowOrigins              []string      `ccf:"allow_origins"`
	AllowMethods              []string      `ccf:"allow_methods"`
	AllowPrivateNetwork       bool          `ccf:"allow_private_network"`
	AllowHeaders              []string      `ccf:"allow_headers"`
	AllowCredentials          bool          `ccf:"allow_credentials"`
	ExposeHeaders             []string      `ccf:"expose_headers"`
	MaxAge                    time.Duration `ccf:"max_age"`
	AllowWildcard             bool          `ccf:"allow_wildcard"`
	AllowBrowserExtensions    bool          `ccf:"allow_browser_extensions"`
	CustomSchemas             []string      `ccf:"custom_schemas"`
	AllowWebSockets           bool          `ccf:"allow_web_sockets"`
	AllowFiles                bool          `ccf:"allow_files"`
	OptionsResponseStatusCode int           `ccf:"options_response_status_code"`
}

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, gin.HandlerFunc]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config Config) (instance gin.HandlerFunc, err error) {
		cfg := cors.Config{
			AllowAllOrigins:           config.AllowAllOrigins,
			AllowOrigins:              config.AllowOrigins,
			AllowMethods:              config.AllowMethods,
			AllowPrivateNetwork:       config.AllowPrivateNetwork,
			AllowHeaders:              config.AllowHeaders,
			AllowCredentials:          config.AllowCredentials,
			ExposeHeaders:             config.ExposeHeaders,
			MaxAge:                    config.MaxAge,
			AllowWildcard:             config.AllowWildcard,
			AllowBrowserExtensions:    config.AllowBrowserExtensions,
			CustomSchemas:             config.CustomSchemas,
			AllowWebSockets:           config.AllowWebSockets,
			AllowFiles:                config.AllowFiles,
			OptionsResponseStatusCode: config.OptionsResponseStatusCode,
		}
		instance = cors.New(cfg)
		return
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
