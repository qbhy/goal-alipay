package alipay

import (
	"github.com/goal-web/supports/logs"
	"github.com/goal-web/supports/utils"
	"github.com/smartwalle/alipay/v3"
	"sync"
)

type factory struct {
	config *Config

	clients sync.Map
}

func (this *factory) Client(name ...string) *alipay.Client {
	var (
		key = utils.DefaultString(name, this.config.Default)

		instance, exists = this.clients.Load(key)
	)
	if exists {
		return instance.(*alipay.Client)
	}

	var config = this.config.Apps[key]
	if config == nil {
		logs.WithField("name", key).Warn("alipay.factory.Client: app is not defined")
		return nil
	}

	var app, err = alipay.New(config.AppId, config.PrivateKey, config.IsProduction)
	if err != nil {
		logs.WithError(err).Error("alipay.factory.Client: create alipay app failed")
		panic(err)
	}

	this.clients.Store(key, app)

	return app
}

func NewFactory(config *Config) Factory {
	return &factory{
		config:  config,
		clients: sync.Map{},
	}
}
