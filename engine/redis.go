package engine

import (
	"strings"

	"github.com/go-redis/redis/v7"
)

// RedisClient is an idiomatic interface for the Redis client,
// adding few methods to interact with the file system.
type RedisClient struct {
	conn *redis.Client
}

// NewRedisClient returns a Redis client
func NewRedisClient(addr, port string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	return &RedisClient{conn: client}, nil
}

// AddFile index a file
func (c *RedisClient) AddFile(key string, content string) error {
	words := strings.Split(content, " ")

	for _, v := range words {
		if err := c.conn.ZAdd(strings.ToLower(v), &redis.Z{
			Score:  float64(1),
			Member: key,
		}).Err(); err != nil {
			return err
		}
	}

	return nil
}

// ZRevRange search for a key
func (c *RedisClient) ZRevRange(key string) ([]string, error) {
	return c.conn.ZRevRange(key, 0, -1).Result()
}

// GetKey returns a key value
func (c *RedisClient) GetKey(key string) (string, error) {
	return c.conn.Get(key).Result()
}

// GetAllKeys returns a key value
func (c *RedisClient) GetAllKeys() (keys []string, err error) {
	keys, _, err = c.conn.Scan(0, "*", 15).Result()

	return
}

// FlushAll drop the database
func (c *RedisClient) FlushAll() error {
	return c.conn.FlushAll().Err()
}

// Close closes the Redis connection
func (c *RedisClient) Close() error {
	return c.conn.Close()
}
