package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/opentracing/opentracing-go"
)

// Rediser ...
type Rediser interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
	Close() error
}

// NewRedisClient ...
func NewRedisClient(ctx context.Context, addr string) *redis.Client {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewRedisClient")
	defer span.Finish()

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

type client struct {
	cl *redis.Client
}

// Set ...
func (c *client) Set(ctx context.Context, key string, value interface{}) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "RedisClient:Set")
	defer span.Finish()
	err := c.cl.Set(key, value, 0).Err()
	return err
}

// Get ...
func (c *client) Get(ctx context.Context, key string) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "RedisClient:Get")
	defer span.Finish()
	val, err := c.cl.Get(key).Result()
	return val, err
}

// Close ...
func (c *client) Close() error {
	return c.cl.Close()
}

// New ...
func New(ctx context.Context, addr string) Rediser {
	r := NewRedisClient(ctx, addr)
	return &client{
		cl: r,
	}
}
