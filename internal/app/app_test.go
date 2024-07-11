package app

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"testing"
	"zg_sql_repo/internal/app/cache"
	"zg_sql_repo/internal/app/kafka"
	"zg_sql_repo/internal/app/repository"
	"zg_sql_repo/internal/app/shard_manager"
)

func TestValidateApp(t *testing.T) {
	err := fx.ValidateApp(
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
	require.NoError(t, err)
}
