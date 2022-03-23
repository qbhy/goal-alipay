package alipay

import "github.com/smartwalle/alipay/v3"

type AppConfig struct {
	AppId      string
	PrivateKey string

	AppPublicCert     string
	AppPublicCertFile string

	AliRootCert     string
	AliRootCertFile string

	AliPublicCert     string
	AliPublicCertFile string

	IsProduction    bool
	OptionFunctions []OptionFun
}

type OptionFun = alipay.OptionFunc

type Config struct {
	Default string
	Apps    map[string]*AppConfig
}
