package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pressly/goose"
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
				"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				db.User, db.Password, db.Host, db.Port, db.DB,
			)
			gdb, err := gorm.Open("mysql", args)

			if err != nil {
				r.Logger.Error("Failed to connect to db", zap.Error(err))
				return
			}
			r.dbs = append(r.dbs, gdb)

			err = r.DoMigrations(db.MigrationsPath, gdb)
			if err != nil {
				r.Logger.Error("Failed to migrate db", zap.Error(err))
				return
			}
		}
	}()
}

func (r *MySQLRepository) Stop(ctx context.Context) {
	r.wg.Wait()
	for _, db := range r.dbs {
		err := db.Close()
		if err != nil {
			r.Logger.Error("Failed to disconnect from db", zap.Error(err))
		}
	}
	r.Logger.Info("Repo started")
}

func (r *MySQLRepository) GetAll(ctx context.Context) ([]*model.Message, error) {

	return nil, nil
}

func (r *MySQLRepository) Create(ctx context.Context, shard int, entity *model.Message) error {
	db := r.dbs[shard]
	return db.Create(entity).Error
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

func (r *MySQLRepository) DoMigrations(pathToMigrations string, db *gorm.DB) error {

	err := goose.Up(db.DB(), pathToMigrations)
	if err != nil {
		if errors.As(err, &goose.ErrNoNextVersion) {
			r.Logger.Info("No migrations to run")
			return nil
		}
		return err
	}

	return nil
}
