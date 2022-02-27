package alipay

import (
	"github.com/goal-web/contracts"
	"github.com/smartwalle/alipay/v3"
)

type ServiceProvider struct {
}

func (s ServiceProvider) Register(app contracts.Application) {
	app.Singleton("alipay", func(config contracts.Config) Factory {
		return NewFactory(config.Get("alipay").(*Config))
	})
	app.Singleton("alipay.client", func(ali Factory) *alipay.Client {
		return ali.Client()
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {
}
