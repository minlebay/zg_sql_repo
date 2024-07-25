package app

import (
	"go.uber.org/fx"
	"zg_sql_repo/internal/app/cache"
	"zg_sql_repo/internal/app/kafka"
	"zg_sql_repo/internal/app/keyvalue_db"
	"zg_sql_repo/internal/app/log"
	"zg_sql_repo/internal/app/repository"
	"zg_sql_repo/internal/app/shard_manager"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			kafka.NewModule(),
			keyvalue_db.NewModule(),
			cache.NewModule(),
			repository.NewModule(),
			shard_manager.NewModule(),
			log.NewModule(),
		),
		fx.Provide(
			NewConfig,
		),
	)
}
