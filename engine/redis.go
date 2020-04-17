package engine

import (
	"strings"

	redis "github.com/go-redis/redis/v7"
)

// NewRedisClient returns a Redis client
func NewRedisClient(addr, port, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// AddFile indexes a file and its words in the database
func AddFile(c *redis.Client, file, content string) error {
	for w, s := range Scan(content) {
		if err := c.ZAdd(file, &redis.Z{
			Score:  float64(s),
			Member: strings.ToLower(w),
		}).Err(); err != nil {
			return err
		}
	}

	return nil
}

// Get search for a key
func Get(c *redis.Client, key string) ([]string, error) {
	return c.ZRevRangeByScore(key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  -1,
	}).Result()
}

// GetWordScoreFromFile retreive the score of a file's word
func GetWordScoreFromFile(c *redis.Client, key, member string) float64 {
	return c.ZScore(key, member).Val()
}

// GetFiles returns all registered files
func GetFiles(c *redis.Client) (keys []string, err error) {
	keys, _, err = c.Scan(0, "*", -1).Result()
	return
}

// FlushAll drops the database
func FlushAll(c *redis.Client) error {
	return c.FlushAll().Err()
}

// Close closes the Redis connection
func Close(c *redis.Client) error {
	return c.Close()
}
