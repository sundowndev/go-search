package engine

import (
	"strings"

	redis "github.com/go-redis/redis/v7"
)

// RedisClient is an idiomatic interface for the Redis client,
// adding few methods to interact with the file system.
type RedisClient struct {
	conn *redis.Client
}

// NewRedisClient returns a Redis client
func NewRedisClient(addr, port, password string, db int) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	return &RedisClient{conn: client}, nil
}

// AddFile index a file
func (c *RedisClient) AddFile(file, content string) error {
	for _, v := range GetWordsFromText(content) {
		if err := c.conn.ZAdd(file, &redis.Z{
			Score:  float64(CountWord(content, v)),
			Member: strings.ToLower(v),
		}).Err(); err != nil {
			return err
		}
	}

	return nil
}

// GetKey search for a key
func (c *RedisClient) GetKey(key string) ([]string, error) {
	return c.conn.ZRevRange(key, 0, -1).Result()
}

// GetScore get score of element
func (c *RedisClient) GetScore(key, member string) float64 {
	return c.conn.ZScore(key, member).Val()
}

// GetAllKeys returns a key value
func (c *RedisClient) GetAllKeys() (keys []string, err error) {
	keys, _, err = c.conn.Scan(0, "*", 35).Result()
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
