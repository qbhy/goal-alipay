package alipay

import "github.com/smartwalle/alipay/v3"

type AppConfig struct {
	AppId           string
	PrivateKey      string
	IsProduction    bool
	OptionFunctions []alipay.OptionFunc
}

type Config struct {
	Default string
	Apps    map[string]*AppConfig
}
