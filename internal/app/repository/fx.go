package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"zg_sql_repo/internal/model"
)

type Repository interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
	GetAll(ctx context.Context) ([]*model.Message, error)
	Create(ctx context.Context, shard int, entity *model.Message) error
	GetById(ctx context.Context, id string) (*model.Message, error)
	Update(ctx context.Context, id string, entity *model.Message) (*model.Message, error)
	Delete(ctx context.Context, id string) error
	GetDbs() []*gorm.DB
	DoMigrations(pathToMigrations string, db *gorm.DB) error
}

func NewModule() fx.Option {

	return fx.Module(
		"repo",
		fx.Provide(
			NewRepositoryConfig,
			fx.Annotate(
				NewMySQLRepository,
				fx.As(new(Repository)),
			),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, r Repository) {
				lc.Append(fx.StartStopHook(r.Start, r.Start))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("repo")
		}),
	)
}
