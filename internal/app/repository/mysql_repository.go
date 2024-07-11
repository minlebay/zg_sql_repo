package repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"sync"
	"zg_sql_repo/internal/model"
)

type MySQLRepository struct {
	Config *Config
	Logger *zap.Logger
	wg     sync.WaitGroup
	dbs    []*gorm.DB
}

func NewMySQLRepository(logger *zap.Logger, config *Config) *MySQLRepository {
	return &MySQLRepository{
		Config: config,
		Logger: logger,
	}
}

func (r *MySQLRepository) Start(ctx context.Context) {
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
			r.dbs = append(r.dbs, gdb)

			// migrations

		}
	}()
}

func (r *MySQLRepository) Stop(ctx context.Context) {
	r.wg.Wait()

	r.Logger.Info("Repo started")
}

func (r *MySQLRepository) GetAll(ctx context.Context) ([]*model.Message, error) {

	return nil, nil
}

func (r *MySQLRepository) Create(ctx context.Context, entity *model.Message) (*model.Message, error) {

	return nil, nil
}

func (r *MySQLRepository) GetById(ctx context.Context, id string) (*model.Message, error) {

	return nil, nil
}

func (r *MySQLRepository) Update(ctx context.Context, id string, entity *model.Message) (*model.Message, error) {

	return nil, nil
}

func (r *MySQLRepository) Delete(ctx context.Context, id string) error {

	return nil
}

func (r *MySQLRepository) GetDbs() []*gorm.DB {
	return r.dbs
}
