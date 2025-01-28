package redis

import (
	"context"
	"time"

	"github.com/EmilioGerdez/Kafka-Logger-Example/pkg/kafka_logger/config"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	client *redis.Client
}

func NewClient(cfg config.RedisConfig) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		// Password: cfg.Password,
		DB: cfg.DB,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic("failed to connect to redis: " + err.Error())
	}

	return &Client{
		client: rdb,
	}
}

// LPush adds a value to the beginning of the list
func (c *Client) LPush(ctx context.Context, key string, values ...interface{}) error {
	return c.client.LPush(ctx, key, values...).Err()
}

// LTrim maintains only last N elements in the list
func (c *Client) LTrim(ctx context.Context, key string, start, stop int64) error {
	return c.client.LTrim(ctx, key, start, stop).Err()
}

// GetRecentLogs retrieves the latest logs
func (c *Client) GetRecentLogs(ctx context.Context, key string, count int64) ([]string, error) {
	return c.client.LRange(ctx, key, 0, count-1).Result()
}

// Close connection
func (c *Client) Close() error {
	return c.client.Close()
}
