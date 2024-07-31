package app

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
	"zg_sql_repo/internal/app/cache"
	"zg_sql_repo/internal/app/kafka"
	"zg_sql_repo/internal/app/keyvalue_db"
	"zg_sql_repo/internal/app/log"
	"zg_sql_repo/internal/app/repository"
	"zg_sql_repo/internal/app/shard_manager"
	"zg_sql_repo/internal/app/tracer"
)

func TestValidateApp(t *testing.T) {
	err := fx.ValidateApp(
		fx.Options(
			kafka.NewModule(),
			keyvalue_db.NewModule(),
			cache.NewModule(),
			repository.NewModule(),
			shard_manager.NewModule(),
			log.NewModule(),
			tracer.NewModule(),
		),
		fx.Provide(
			NewConfig,
		),
	)
	require.NoError(t, err)
}
