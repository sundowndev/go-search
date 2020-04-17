package engine

import (
	"log"
	"testing"

	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
)

const (
	addr     = "0.0.0.0"
	port     = "6379"
	password = ""
	DB       = 0
)

func TestRedisClient(t *testing.T) {
	assert := assert.New(t)

	client := redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: password, // no password set
		DB:       DB,       // use default DB
	})
	defer client.FlushAll()
	defer client.Close()

	redisClient, err := NewRedisClient(addr, port, password, DB)
	if err != nil {
		log.Fatal(err)
	}
	defer redisClient.Close()

	t.Run("RedisClient", func(t *testing.T) {
		t.Run("should use ZADD", func(t *testing.T) {
			client.FlushAll()

			err := redisClient.AddFile("file1", "word")
			assert.Equal(nil, err, "should be equal")

			err = redisClient.AddFile("file2", "word")
			assert.Equal(nil, err, "should be equal")

			values, err2 := client.ZRevRange("file1", 0, -1).Result()
			assert.Equal(nil, err2, "should be equal")

			assert.Equal([]string{"word"}, values, "should be equal")
		})

		t.Run("should use ZREVRANGE", func(t *testing.T) {
			client.FlushAll()

			err := client.ZAdd("word", &redis.Z{
				Score:  float64(1),
				Member: "file",
			}).Err()
			assert.Equal(nil, err, "should be equal")

			values, err2 := redisClient.Get("word")
			assert.Equal(nil, err2, "should be equal")

			assert.Equal([]string{"file"}, values, "should be equal")
		})

		t.Run("should use FLIUSHALL", func(t *testing.T) {
			client.FlushAll()

			err := client.Set("test", "value", 0).Err()
			assert.Equal(nil, err, "should be equal")

			err = redisClient.FlushAll()
			assert.Equal(nil, err, "should be equal")

			err = client.Get("test").Err()
			assert.Equal(redis.Nil, err, "should be equal")
		})
	})
}
