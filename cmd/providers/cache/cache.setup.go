package cache

import (
	"fmt"
	"runtime/debug"
	"synergize/backend-test/pkg/entity"
	"synergize/backend-test/pkg/facades"
	"synergize/backend-test/pkg/helper"

	"github.com/go-redis/redis"
)

func bootCache() entity.Cache {
	config := facades.Config.Get("database.redis").(map[string]any)

	cache := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config["host"], config["port"]),
		DB:       config["database"].(int),
		Password: config["password"].(string),
	})

	ping, err := cache.Ping().Result()
	if err != nil {
		err := fmt.Errorf("failed to link redis:%s, %s\n%+v", ping, err, string(debug.Stack()))
		fmt.Println(err.Error())
	}

	helper.PanicIfError(err)

	return &Redis{
		Redis: cache,
	}
}
