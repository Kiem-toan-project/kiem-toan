package httpreq

import (
	"net/http"
)

type RestyConfig struct {
	Client *http.Client
}

type Resty struct {
	resty.Client
}

func NewResty(cfg RestyConfig) *Resty {
	client := &Resty{}
	if cfg.Client == nil {
		client.Client = *resty.New()
	} else {
		client.Client = *resty.NewWithClient(cfg.Client)
	}
	return client
}
