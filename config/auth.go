package config

import (
	"github.com/caarlos0/env"
)

type Authentication struct {
	VerifyKey string `env:"API_AUTHENTICATION_SECRET,required"`
	Algorithm string `env:"API_AUTHENTICATION_ALGORITHM" envDefault:"HS256"`
}

func (c *ConfigImpl) JWT() *Authentication {
	if c.jwt != nil {
		return c.jwt
	}

	c.Lock()
	defer c.Unlock()

	jwt := &Authentication{}
	if err := env.Parse(jwt); err != nil {
		panic(err)
	}

	c.jwt = jwt

	return c.jwt
}
