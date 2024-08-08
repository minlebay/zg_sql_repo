package shard_manager

import (
	"context"
	"encoding/json"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"hash/crc32"
	"strconv"
	"sync"
	"zg_sql_repo/internal/app/cache"
	"zg_sql_repo/internal/app/keyvalue_db"
	"zg_sql_repo/internal/app/repository"
	"zg_sql_repo/internal/model"
)

type Manager struct {
	Config     *Config
	Logger     *zap.Logger
	Cache      cache.Cache
	KeyValueDB keyvalue_db.KValueDB
	Repository repository.Repository
	Tracer     trace.Tracer
	wg         sync.WaitGroup
}

func NewManager(
	logger *zap.Logger,
	config *Config,
	cache cache.Cache,
	kvdb keyvalue_db.KValueDB,
	repo repository.Repository,
) *Manager {
	tracer := otel.Tracer("sql_shard-manager")
	return &Manager{
		Config:     config,
		Logger:     logger,
		Cache:      cache,
		KeyValueDB: kvdb,
		Repository: repo,
		Tracer:     tracer,
	}
}

func (m *Manager) StartManager() {
	m.Logger.Info("Shard manager started")
}

func (m *Manager) StopManager() {
	m.wg.Wait()
	m.Logger.Info("Shard manager stopped")
}

func (m *Manager) Consume(ctx context.Context, msg *model.Message) error {
	m.wg.Add(1)
	defer m.wg.Done()

	ctx, span := m.Tracer.Start(ctx, "Consume message")
	defer span.End()

	shardIndex, err := m.GetShardIndex(ctx, msg.Uuid)
	if err != nil {
		m.Logger.Error("Failed to get shard index", zap.Error(err))
		return err
	}
	span.SetAttributes(attribute.Int("shard_index", shardIndex))

	err = m.Repository.Create(ctx, shardIndex, msg)
	if err != nil {
		m.Logger.Error("Failed to store message", zap.Error(err))
		return err
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		m.Logger.Error("Failed to marshal message", zap.Error(err))
		return err
	}
	err = m.Cache.Put(msg.Uuid, bytes)
	if err != nil {
		m.Logger.Error("Message stored to db, but failed to cache message", zap.Error(err))
	}

	m.Logger.Info("Message stored and cached", zap.String("uuid", msg.Uuid), zap.Int("shard", shardIndex))

	bytes = []byte(strconv.Itoa(shardIndex))
	err = m.KeyValueDB.Put(msg.Uuid, bytes)
	if err != nil {
		m.Logger.Error("Failed to store index", zap.Error(err))
		return err
	}
	m.Logger.Info("Index stored", zap.String("uuid", msg.Uuid), zap.Int("shard", shardIndex))
	return nil
}

func (m *Manager) GetShardIndex(ctx context.Context, uuid string) (int, error) {
	dbsCount := len(m.Repository.GetDbs())
	if dbsCount == 0 {
		return 0, nil
	}
	uuidBytes := []byte(uuid)
	hash := crc32.ChecksumIEEE(uuidBytes)
	shardNumber := int(hash) % dbsCount
	return shardNumber, nil
}
