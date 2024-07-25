package cache

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var _ Cache = (*Redis)(nil)

type Cache interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
	Get(key string) (out []byte, err error)
	Put(key string, value []byte) (err error)
	Delete(key string) (err error)
	Iterate(filter string) (keys []string, err error)
}

func NewModule() fx.Option {

	return fx.Module(
		"cache",
		fx.Provide(
			NewCacheConfig,
			fx.Annotate(
				NewRedis,
				fx.As(new(Cache)),
			),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, c Cache) {
				lc.Append(fx.StartStopHook(c.Start, c.Stop))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("sql_cache")
		}),
	)
}
