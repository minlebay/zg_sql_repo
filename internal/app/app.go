package app

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"zg_sql_repo/internal/app/cache"
	"zg_sql_repo/internal/app/kafka"
	"zg_sql_repo/internal/app/repository"
	"zg_sql_repo/internal/app/shard_manager"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			kafka.NewModule(),
			cache.NewModule(),
			repository.NewModule(),
			shard_manager.NewModule(),
		),
		fx.Provide(
			zap.NewProduction,
			NewConfig,
		),
	)
}
