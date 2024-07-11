package shard_manager

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"hash/crc32"
	"strconv"
	"sync"
	"zg_sql_repo/internal/app/cache"
	"zg_sql_repo/internal/app/repository"
	"zg_sql_repo/internal/model"
)

type Manager struct {
	Config     *Config
	Logger     *zap.Logger
	Cache      cache.Cache
	Repository repository.Repository
	wg         sync.WaitGroup
}

func NewManager(
	logger *zap.Logger,
	config *Config,
	cache cache.Cache,
	repo repository.Repository,
) *Manager {
	return &Manager{
		Config:     config,
		Logger:     logger,
		Cache:      cache,
		Repository: repo,
	}
}

func (m *Manager) StartManager(ctx context.Context) {
	m.Logger.Info("Shard manager started")
}

func (m *Manager) StopManager(ctx context.Context) {
	m.wg.Wait()
	m.Logger.Info("Shard manager stopped")
}

func (m *Manager) Consume(ctx context.Context, msg *model.Message) error {
	m.wg.Add(1)
	defer m.wg.Done()

	// Calculate shard number
	shardIndex, err := m.GetShardIndex(ctx, msg.Uuid)
	if err != nil {
		m.Logger.Error("Failed to get shard index", zap.Error(err))
		return err
	}

	created, err := m.Repository.Create(ctx, msg)
	if err != nil {
		m.Logger.Error("Failed to store message", zap.Error(err))
		return err
	}

	bytes, err := json.Marshal(created)
	if err != nil {
		m.Logger.Error("Failed to marshal message", zap.Error(err))
		return err
	}
	m.Cache.Put(created.Uuid, bytes)

	m.Logger.Info("Message stored and cached", zap.String("uuid", created.Uuid), zap.Int("shard", shardIndex))

	bytes = []byte(strconv.Itoa(shardIndex))
	err = m.Cache.Put(created.Uuid, bytes)
	if err != nil {
		m.Logger.Error("Failed to store index", zap.Error(err))
		return err
	}
	m.Logger.Info("Index stored", zap.String("uuid", created.Uuid), zap.Int("shard", shardIndex))
	return nil
}

func (m *Manager) GetShardIndex(ctx context.Context, uuid string) (int, error) {
	uuidBytes := []byte(uuid)
	hash := crc32.ChecksumIEEE(uuidBytes)
	shardNumber := int(hash) % len(m.Repository.GetDbs())
	return shardNumber, nil
}
