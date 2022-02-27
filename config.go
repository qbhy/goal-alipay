package alipay

type AppConfig struct {
	AppId        string
	PrivateKey   string
	IsProduction bool
}

type Config struct {
	Default string
	Apps    map[string]*AppConfig
}
