package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"sync"
	"zg_sql_repo/internal/model"
)

type Repository struct {
	Config           *Config
	Logger           *zap.Logger
	wg               sync.WaitGroup
	DBs              []*any
	CancelFunc       context.CancelFunc
	ClientDisconnect func()
}

func NewRepository(logger *zap.Logger, config *Config) *Repository {
	return &Repository{
		Config: config,
		Logger: logger,
	}
}

func (r *Repository) StartRepository(ctx context.Context) {
	go func() {
		for _, db := range r.Config.Dbs {
			_ = db
		}
	}()
}

func (r *Repository) StopRepository(ctx context.Context) {
	r.wg.Wait()
	r.ClientDisconnect()
	r.CancelFunc()

	r.Logger.Info("Repo started")
}

func (r *Repository) GetAll(ctx context.Context, db mongo.Database) ([]*model.Message, error) {

	return nil, nil
}

func (r *Repository) Create(ctx context.Context, entity *model.Message) (*model.Message, error) {

	return nil, nil
}

func (r *Repository) GetById(ctx context.Context, id string) (*model.Message, error) {

	return nil, nil
}

func (r *Repository) Update(ctx context.Context, id string, entity *model.Message) (*model.Message, error) {

	return nil, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {

	return nil
}

func (r *Repository) GetConnected() {

}
