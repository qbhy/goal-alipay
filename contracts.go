package alipay

import "github.com/smartwalle/alipay/v3"

type Factory interface {
	Client(name ...string) *alipay.Client
}
