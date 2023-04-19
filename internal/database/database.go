package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

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

func GetChapter(aid int, volume int, chapter int) (string, bool) {
	result, err := Redis.Get(Ctx, fmt.Sprintf("%d-%d-%d", aid, volume, chapter)).Result()
	util.ErrorHandler(err, false)
	if err != nil {
		return result, false
	}
	return result, true
}

func AddChapter(aid int, volume int, chapter int, content string) {
	err := Redis.Set(Ctx, fmt.Sprintf("%d-%d-%d", aid, volume, chapter), content, time.Hour*24).Err()
	util.ErrorHandler(err, false)
}
