package config

import (
	"github.com/caarlos0/env"
)

type Bot struct {
	Token string `env:"API_BOT_TOKEN,required"`
}

func (c *ConfigImpl) Bot() *Bot {
	if c.bot != nil {
		return c.bot
	}

	c.Lock()
	defer c.Unlock()

	var b Bot
	if err := env.Parse(&b); err != nil {
		panic(err)
	}

	c.bot = &b
	return c.bot
}
