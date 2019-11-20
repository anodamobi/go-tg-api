package config

import (
	"github.com/caarlos0/env"
)

type OneSignal struct {
	AppID  string `env:"API_ONE_SIGNAL_APP_ID"`
	APIKey string `env:"API_ONE_SIGNAL_API_KEY"`
}

func (c *ConfigImpl) OneSignal() *OneSignal {
	if c.oneSignal != nil {
		return c.oneSignal
	}

	c.Lock()
	defer c.Unlock()

	oneSignal := &OneSignal{}
	if err := env.Parse(oneSignal); err != nil {
		panic(err)
	}

	c.oneSignal = oneSignal

	return c.oneSignal
}
