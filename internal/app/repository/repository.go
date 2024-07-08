package repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"sync"
	"zg_sql_repo/internal/model"
)

type Repository struct {
	Config *Config
	Logger *zap.Logger
	wg     sync.WaitGroup
	DBs    []*gorm.DB
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
			args := fmt.Sprintf(
				"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
				db.User, db.Password, db.DB,
			)
			gdb, err := gorm.Open("mysql", args)
			if err != nil {
				r.Logger.Error("Failed to connect to db", zap.Error(err))
				return
			}
			r.DBs = append(r.DBs, gdb)

			// migrations

		}
	}()
}

func (r *Repository) StopRepository(ctx context.Context) {
	r.wg.Wait()

	r.Logger.Info("Repo started")
}

func (r *Repository) GetAll(ctx context.Context) ([]*model.Message, error) {

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
