package cache

import (
	"context"
	"go.uber.org/zap"
)

type Stub struct {
	Config *Config
	Logger *zap.Logger
}

func NewCacheStub(config *Config, logger *zap.Logger) *Stub {
	return &Stub{
		Config: config,
		Logger: logger,
	}
}

func (c *Stub) Start(ctx context.Context) {

}

func (c *Stub) Stop(ctx context.Context) {

}

func (c *Stub) Get(key string) (out []byte, err error) {
	return []byte("stub"), nil
}

func (c *Stub) Put(key string, value []byte) (err error) {
	return nil
}

func (c *Stub) Iterate(filter string) (keys []string, err error) {
	return []string{"stub_key"}, nil
}

func (c *Stub) Delete(key string) (err error) {
	return nil
}
