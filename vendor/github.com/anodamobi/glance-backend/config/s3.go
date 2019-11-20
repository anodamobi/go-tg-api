package config

import (
	"github.com/caarlos0/env"
	"github.com/minio/minio-go"
)

type S3 struct {
	AccessKey string `env:"API_S3_ACCESS_KEY"`
	SecretKey string `env:"API_S3_SECRET_KEY"`
	Endpoint  string `env:"API_S3_ENDPOINT"`
}

func (c *ConfigImpl) S3() *minio.Client {
	if c.s3 != nil {
		return c.s3
	}

	c.Lock()
	defer c.Unlock()

	s3 := &S3{}
	if err := env.Parse(s3); err != nil {
		panic(err)
	}

	client, err := minio.New(s3.Endpoint, s3.AccessKey, s3.SecretKey, true)
	if err != nil {
		c.log.WithError(err).Error("failed to create minio default client")
		panic(err)
	}

	c.s3 = client

	return client
}
