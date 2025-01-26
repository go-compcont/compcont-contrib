package ddddocr

import (
	"errors"
	"net/url"
	"time"

	"github.com/go-compcont/compcont-core"
	"github.com/go-resty/resty/v2"
)

const TypeID compcont.ComponentTypeID = "contrib.ddddocr"

type Config struct {
	BaseURL string `ccf:"base_url"`
}

func buildComponent(cfg Config) (c OCR, err error) {
	if cfg.BaseURL == "" {
		err = errors.New("url is required")
		return
	}

	_, err = url.Parse(cfg.BaseURL)
	if err != nil {
		return
	}

	restyCli := resty.New().SetTimeout(3 * time.Second)

	ocr := &DdddOCR{
		url:    cfg.BaseURL,
		client: restyCli,
	}
	c = ocr
	return
}

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, any]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config Config) (instance any, err error) {
		return buildComponent(config)
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
