package cache

import (
	"encoding/json"
	"fmt"
	"service/models"
	"time"

	"github.com/go-redis/redis"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) CovidCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cache.host,
		Password: "0MEbNW8gM3zDoaKtDWwjy6WfkQK5zJ9X",
		DB: cache.db,
	})
}

func (cache *redisCache) Set(key *models.Location, value *models.Response) {
	client := cache.GetClient()
	stringKey, err := json.Marshal(key)
	if err != nil {
		fmt.Println(err)
	}
	data, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}
	client.Set(string(stringKey), data, cache.expires*time.Second)
	fmt.Println("data set to cache")
}
func (cache *redisCache) Get(key *models.Location) *models.Response {
	client := cache.GetClient()
	stringKey, err := json.Marshal(key)
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get(string(stringKey)).Result()
	if err != nil {
		fmt.Println(err)
	}
	result := models.Response{}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("data from cache")
	return &result
}
