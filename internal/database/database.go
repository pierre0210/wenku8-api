package database

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/pierre0210/wenku8-api/internal/util"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client
var Ctx context.Context

func InitRedis() {
	var redisAddr string
	var redisPort string
	var redisDB int
	redisPass := os.Getenv("REDIS_PASS")

	if os.Getenv("REDIS_ADDR") == "" {
		redisAddr = "localhost"
	} else {
		redisAddr = os.Getenv("REDIS_ADDR")
	}

	if os.Getenv("REDIS_PORT") == "" {
		redisPort = "6379"
	} else {
		redisPort = os.Getenv("REDIS_PORT")
	}

	if os.Getenv("REDIS_DB") == "" {
		redisDB = 0
	} else {
		redisDB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	}

	Ctx = context.Background()

	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisAddr, redisPort),
		Password: redisPass,
		DB:       redisDB,
	})

	_, err := Redis.Ping(Ctx).Result()
	util.ErrorHandler(err, true)
}
