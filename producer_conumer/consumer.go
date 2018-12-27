package producer_conumer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var (
	topic = "test02"
)

type Client struct {

	Producer 	*Producer
	Consumer 	*Consumer
}

type Consumer struct {

}

func (consumer *Consumer) Consuming(wg *sync.WaitGroup, brokers, topics []string, groupId string) {
	defer wg.Done()
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_0
	//新建消费者
	cons, err := sarama.NewConsumer([]string{"39.108.171.31:9092"}, config)
	if err != nil {
		fmt.Println("sarama new consumer error. ", err.Error())
		return
	}
	defer cons.Close()

	part, err := cons.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("consumer partition error. ", err.Error())

	}
	defer part.Close()

	for {
		select {
		case msg := <- part.Messages():
			fmt.Println("msg offset: ", msg.Offset, " partition:", msg.Partition, " timestamp: ", msg.Timestamp.Format("2006-Jan-02 15:04"), " value: ", string(msg.Value))
		case err := <- part.Errors():
			fmt.Println("consume message error. ", err.Error())
		}
	}
}

