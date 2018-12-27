package producer_conumer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"time"
)

type Producer struct {

}

func (producer * Producer) Producing() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	prod, err := sarama.NewSyncProducer([]string{"39.108.171.31:9092"}, config)
	if err != nil {
		fmt.Println("sarama.NewSyncProducer err, message=", err.Error())
		return
	}
	defer prod.Close()

	semiValue := "this is a message produced at %d"
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf(semiValue, i)
		producedMessage := &sarama.ProducerMessage{
			Topic: 	topic,
			Value: 	sarama.ByteEncoder(value),
		}
		part, offset, err := prod.SendMessage(producedMessage)
		if err != nil {
			fmt.Printf("send message(%s) err=%s \n", value, err.Error())

		} else {
			fmt.Fprintf(os.Stdout, value + "Send Success! partition=%d, offset=%d \n", part, offset)
		}
		time.Sleep(2*time.Second)

	}
}