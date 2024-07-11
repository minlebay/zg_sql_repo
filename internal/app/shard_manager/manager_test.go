package shard_manager

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
	"zg_sql_repo/internal/app/cache"
	"zg_sql_repo/internal/app/repository"
	"zg_sql_repo/internal/model"
)

// MockLogger is a mock type for the Logger
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockLogger) Error(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func TestManager_Consume(t *testing.T) {
	cacheStub := cache.NewCacheStub(nil, nil)
	repoStub := repository.NewRepositoryStub(nil, nil)
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	m := NewManager(logger, nil, cacheStub, repoStub)

	msg := &model.Message{
		Uuid: uuid.New().String(),
	}

	err = m.Consume(nil, msg)

	require.NoError(t, err)
}

func BenchmarkManager_Consume1000(b *testing.B) {
	cacheStub := cache.NewCacheStub(nil, nil)
	repoStub := repository.NewRepositoryStub(nil, nil)
	logger, _ := zap.NewDevelopment()

	m := NewManager(logger, nil, cacheStub, repoStub)

	msg := &model.Message{
		Uuid: uuid.New().String(),
	}

	for i := 0; i < b.N; i++ {
		_ = m.Consume(nil, msg)
	}
}
