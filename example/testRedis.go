package main

import (
	"fmt"
	"pdriver/testRedis"
)

func main() {

	client, err := testRedis.NewCacheSentinel()
	if err != nil {
		fmt.Println("new Client error.", err.Error())
		return
	}
	fmt.Println("check option")
	ret, _ := client.Get("say").Result()
	fmt.Println("test result: ", ret)

	testRedis.TestClient()

	//testRedis.ConnSentinel()

	//client1 := redis.NewClient(&redis.Options{
	//	Addr:
	//})

}


//
//// 创建 redis 客户端
//func createClient() *redis.Client {
//
//	client := redis.NewClient(&redis.Options{
//		Addr:     "172.27.119.38:6379",
//		Password: "",
//		DB:       0,
//	})
//
//	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
//	//pong, err := client.Ping().Result()
//	//fmt.Println(pong, err)
//
//	return client
//}