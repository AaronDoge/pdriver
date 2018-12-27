package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {

	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: 	"redis-master",
		SentinelAddrs: []string{"172.27.119.38:27000", "172.27.119.39:27000", "172.27.119.40:27000", "172.27.119.41:27000"},
		Password: 	"",
		PoolSize: 	4,
	})

	fmt.Println(">>>check password:",client.Options().Password)

	err := client.Set("nihao", "bonjour", 0).Err()
	if err != nil {
		fmt.Println("redis set failed. ", err.Error())
		return
	}

	startTime := time.Now()
	val, err := client.Get("nihao").Result()
	if err != nil {
		fmt.Println("redis get failed. ", err.Error())
		return
	}
	fmt.Println("value is", val)
}
