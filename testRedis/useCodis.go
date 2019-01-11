package testRedis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func initCodis(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: 	addr,
		PoolSize: 500,
		PoolTimeout: 90 * time.Second,
		MaxRetries: 3,
		IdleTimeout: 90 * time.Second,
	})

	return client
}

// 订阅
func SubscribeCase() {
	client := initCodis("192.168.32.129:6379")
	pubSub := client.Subscribe("test")

	for {
		//time.Sleep(3)
		recMsg, err := pubSub.Receive()
		//msg, err := pubSub.ReceiveMessage()
		//msgStr := msg.String()

		if err != nil {
			fmt.Println("订阅错误", err.Error())
			return
		}
		fmt.Println(">>>check msg", recMsg)

		fmt.Println(recMsg)
		//switch recMsg.(type) {
		//case string:
		//	fmt.Println("subscribe message is:", recMsg)
		//case error:
		//	fmt.Println("错误")
		//}
	}

}

func Get(key string) string {
	addr := "192.168.32.129:19000"
	client := initCodis(addr)

	val, err := client.Get(key).Result()
	if err != nil {
		fmt.Println("get value error", err.Error())
		return ""
	}

	return val
}

func Set(key, val string) error {
	addr := "192.168.32.128:6379"
	client := initCodis(addr)

	err := client.Set(key, val, 0).Err()
	if err != nil {
		fmt.Println("set value error", err.Error())
		return err
	}

	return nil
}

func InitCodisProxyCluster() *redis.ClusterClient {
	addrs := []string{"192.168.32.128:19000","192.168.32.129:19000","192.168.32.130:19000"}
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
		PoolSize: 300,
		PoolTimeout: 90 * time.Second,
		DialTimeout: 15 * time.Second,
		MaxRetries: 3,

	})

	return clusterClient
}

func ClusterGet(key string, client *redis.ClusterClient) string {
	//client := initCodisProxyCluster()
	val, err := client.Get(key).Result()
	if err != nil {
		fmt.Println("cluster get value error", err.Error())
		return ""
	}
	return val
}

func ClusterSet(key, val string, client *redis.ClusterClient) error {
	//client := initCodisProxyCluster()
	if err := client.Set(key, val, 0).Err(); err != nil {
		fmt.Println("cluster set error", err.Error())
		return err
	}
	return nil
}