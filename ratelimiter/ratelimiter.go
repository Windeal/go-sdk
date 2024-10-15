package ratelimiter

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

// RateLimiter struct
type RateLimiter struct {
	h   *log.Helper
	rdb *redis.Client
}

// NewRateLimiter creates a new RateLimiter
func NewRateLimiter(logger log.Logger, rdb *redis.Client) *RateLimiter {
	rl := &RateLimiter{
		h:   log.NewHelper(logger),
		rdb: rdb,
	}

	return rl
}

// Allow checks if a request is allowed under the rate limit
func (r *RateLimiter) Allow(ctx context.Context, key string, window time.Duration, maxRequests int) (bool, error) {
	now := time.Now().UnixNano()

	// Use a transaction to ensure atomicity
	_, err := r.rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.ZAdd(ctx, key, &redis.Z{Score: float64(now), Member: now})
		pipe.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", now-int64(window)))
		pipe.ZCard(ctx, key)
		pipe.Expire(ctx, key, window)
		return nil
	})
	if err != nil {
		r.h.WithContext(ctx).Errorf("r.data.rdb.TxPipelined error, %+v", err)
		return false, err
	}

	count, err := r.data.rdb.ZCard(ctx, key).Result()
	if err != nil {
		r.h.WithContext(ctx).Errorf("r.data.rdb.TxPipelined error, %+v", err)
		return false, err
	}

	return count <= int64(maxRequests), nil
}
