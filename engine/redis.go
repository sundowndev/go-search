package engine

import (
	"encoding/json"
	"io/ioutil"
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
func (c *RedisClient) AddFile(file string, score int) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	words := strings.Split(string(data), " ")

	wordsJSONArray, err := json.Marshal(&words)
	if err != nil {
		return err
	}

	return c.conn.Set(file, wordsJSONArray, 0).Err()
}

// GetKey returns a key value
func (c *RedisClient) GetKey(key string) (string, error) {
	// Utilisez Redigo pour lire toutes les valeurs de la clef, et les
	// placer dans une tranche de chaînes. Renvoyez une erreur si nécessaire.
	return c.conn.Get(key).Result()
}

// GetAllKeys returns a key value
func (c *RedisClient) GetAllKeys() (keys []string, err error) {
	// Utilisez Redigo pour lire toutes les valeurs de la clef, et les
	// placer dans une tranche de chaînes. Renvoyez une erreur si nécessaire.
	keys, _, err = c.conn.Scan(0, "*", 15).Result()

	return
}

// Close closes the Redis connection
func (c *RedisClient) Close() error {
	return c.conn.Close()
}
