package cache

import (
	"context"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"sort"
	"strconv"
	"sync"
	"time"
)

type Redis struct {
	Config  *Config
	Logger  *zap.Logger
	db      *redis.Client
	wg      sync.WaitGroup
	expires time.Duration
}

func NewRedis(logger *zap.Logger, config *Config) *Redis {
	return &Redis{
		Config: config,
		Logger: logger,
	}
}

func (r *Redis) Start(ctx context.Context) {
	go func() {
		numdb, err := strconv.ParseInt(r.Config.DB, 10, 64)
		if err != nil {
			r.Logger.Error("Failed to parse DB", zap.Error(err))
		}

		r.expires, err = time.ParseDuration(r.Config.ExpTime)
		if err != nil {
			r.Logger.Error("Failed to parse expiration time", zap.Error(err))
		}

		r.db = redis.NewClient(&redis.Options{
			Addr: r.Config.Address,
			DB:   int(numdb),
		})
	}()
}

func (r *Redis) Stop(ctx context.Context) {
	r.wg.Wait()
	err := r.db.Close()
	if err != nil {
		r.Logger.Error("Failed to disconnect from Redis", zap.Error(err))
	}
}

func (r *Redis) Get(key string) (out []byte, err error) {
	out, err = r.db.Get(key).Bytes()
	return
}

func (r *Redis) Put(key string, value []byte) (err error) {
	err = r.db.Set(key, string(value), r.expires).Err()
	return
}

func (r *Redis) Delete(key string) (err error) {
	err = r.db.Del(key).Err()
	return
}

func (r *Redis) Iterate(filter string) (out []string, err error) {
	if filter != "" {
		filter += "*"
	}
	iter := r.db.Scan(0, filter, 0).Iterator()
	for iter.Next() {
		key := iter.Val()
		out = append(out, key)
	}
	sort.Strings(out)
	err = iter.Err()
	return
}
