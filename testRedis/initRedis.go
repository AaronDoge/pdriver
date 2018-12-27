package testRedis

import (
	"fmt"
	"github.com/go-redis/redis"
	"pdriver/etc"
	"time"
)

type Server struct {
	Cfg 	*Config
	SentinelClient 	*redis.SentinelClient
	SentinelAddrs 	[]*redis.Client
	CurrMasterAddr	[]*redis.Client

}

type Config struct {
	MasterName 	string
}



func NewCacheSentinel() (*redis.Client,error) {

	var c = &etc.Config{}
	if err := c.LoadConfig("etc/configuration.json"); err != nil {
		fmt.Println("load config error. ", err.Error())
		return nil, err
	}

	opt := &redis.FailoverOptions{
		MasterName: 	c.Redis.MasterName,
		SentinelAddrs: 	c.Redis.Sentinel,
		Password: 		c.Redis.Password,
		PoolSize: 		c.Redis.PoolSize,

		PoolTimeout: 	time.Hour,
		DialTimeout: 	time.Second,
		IdleTimeout:	time.Millisecond,
		IdleCheckFrequency: time.Millisecond,

	}
	client := redis.NewFailoverClient(opt)

	return client, nil
}

func (s *Server) initSentinelClient() {
	var c = &etc.Config{}
	if err := c.LoadConfig("etc/configuration.json"); err != nil {
		fmt.Println("load config error. ", err.Error())
		return
	}

	var sentinelClient = redis.NewSentinelClient(&redis.Options{
		//MasterName: 	c.Redis.MasterName,
		//SentinelAddrs: 	c.Redis.Sentinel,
		Addr: 			c.Redis.Host,
		Password: 		c.Redis.Password,
		PoolSize: 		c.Redis.PoolSize,

		PoolTimeout: 	time.Hour,
		DialTimeout: 	time.Second,
		IdleTimeout:	time.Millisecond,
		IdleCheckFrequency: time.Millisecond,
	})

	s.SentinelClient = sentinelClient
}

func (s *Server)getSentinelClient() {


}

//
func TestClient() {

	var c = &etc.Config{}
	if err := c.LoadConfig("etc/configuration.json"); err != nil {
		fmt.Println("load config error. ", err.Error())
		//return nil, err
	}

	var sentinelClient = redis.NewSentinelClient(&redis.Options{
		//MasterName: 	c.Redis.MasterName,
		//SentinelAddrs: 	c.Redis.Sentinel,
		Addr: 			c.Redis.Host,
		Password: 		c.Redis.Password,
		PoolSize: 		c.Redis.PoolSize,

		PoolTimeout: 	time.Hour,
		DialTimeout: 	time.Second,
		IdleTimeout:	time.Millisecond,
		IdleCheckFrequency: time.Millisecond,
	})


	master, err := sentinelClient.GetMasterAddrByName("redis-master").Result()
	if err != nil {
		fmt.Println("get master name failed", err.Error())
	}
	fmt.Println("master name: ", master)
	sentinels, err := sentinelClient.Sentinels("redis-master").Result()

	fmt.Println("check sentinels num:", len(sentinels))
	if err != nil {
		fmt.Println("get sentinels failed", err.Error())
	}
	for i := 0; i < len(sentinels); i++ {
		fmt.Println("sentinels: ", sentinels[i])
		//fmt.Println("sentinels type:", reflect.TypeOf(sentinels[0]))
	}

	var sentinelAddrs []string

	for _, sentinel := range sentinels {
		vals := sentinel.([]interface{})
		for i := 0; i < len(vals); i += 2 {
			key := vals[i].(string)
			if key == "name" {
				sentinelAddr := vals[i+1].(string)
				fmt.Println("check sentinel addr:", sentinelAddr)
				//if !contains(c.sentinelAddrs, sentinelAddr) {
				//	internal.Logf("sentinel: discovered new sentinel=%q for master=%q",
				//		sentinelAddr, c.masterName)
				//	c.sentinelAddrs = append(c.sentinelAddrs, sentinelAddr)
				//}
				sentinelAddrs = append(sentinelAddrs, sentinelAddr)
			}
		}
	}
	fmt.Println("check final:", sentinelAddrs)

	var clients []*redis.Client
	for i := 0; i < len(sentinelAddrs); i++ {
		client1 := redis.NewClient(&redis.Options{
			Addr: 	sentinelAddrs[i],
			MaxRetries: 	3,
			PoolSize: 		1000,
			PoolTimeout: 	90 * time.Second,
			DialTimeout: 	2 * time.Second,

		})
		clients = append(clients, client1)
	}

}

func ConnSentinel() {
	sentinels := []string{"172.27.119.39:6379", "172.27.119.40:6379", "172.27.119.38:6379"}

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:	sentinels,
		MaxRetries: 	3,
		MaxRedirects: 	3,
		Password: 	"",
		PoolSize: 	100,
		PoolTimeout: 	90 * time.Second,
		DialTimeout: 	2 * time.Second,
	})

	val, err := client.Get("say").Result()
	if err != nil {
		fmt.Println("get say error ", err.Error())
	}

	fmt.Println("get val:", val)
}
