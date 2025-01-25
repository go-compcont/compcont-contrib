package ddddocr

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/go-resty/resty/v2"
)

type OCR interface {
	OCR(io.Reader) (s string, err error)
}

type DdddOCR struct {
	url    string
	client *resty.Client
}

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func New(url string, client *resty.Client) OCR {
	return &DdddOCR{
		url:    url,
		client: client,
	}
}

func (d *DdddOCR) OCR(r io.Reader) (s string, err error) {
	resp, err := d.client.R().
		SetFileReader("file", "image.png", r).
		Post(d.url + "/ocr")
	if err != nil {
		return
	}
	var res APIResponse
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return
	}
	if res.Code != 200 {
		err = errors.New(res.Message)
		return
	}
	s = res.Data
	return
}
