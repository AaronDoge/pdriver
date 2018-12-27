package main

import (
	"pdriver/producer_conumer"
)

func main() {
	client := producer_conumer.Client{
		Producer: &producer_conumer.Producer{},
		Consumer: &producer_conumer.Consumer{},
	}

	client.Producer.Producing()

	// wg *sync.WaitGroup, brokers, topics []string, groupId string
	//var wg *sync.WaitGroup
	//topics := []string{"test02"}
	//brokers := []string{""}
	//groupId := ""
	//client.Consumer.Consuming(wg, brokers, topics, groupId)
}


