package testRedis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func Get(k string) string {
	codis := initCodis()

	val, err := codis.Get(k).Result()
	if err != nil {
		fmt.Println("get value error.", err.Error())
		return ""
	}
	return val
}

func Set(key, val string) error {
	codis := initCodis()

	err := codis.Set(key, val, 0).Err()
	if err != nil {
		fmt.Println("set key/value error.", err.Error())
		return err
	}

	return nil
}

func initCodis() *redis.Client  {

	client := redis.NewClient(&redis.Options{
		Addr: 	"39.108.171.31:6379",
		Password: "",
		PoolTimeout: 90 * time.Second,
		PoolSize: 	10,
		DB: 	0,
	})

	return client
}


