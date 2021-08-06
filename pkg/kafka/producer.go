package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
)

func Produce(topic string, message string) {
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	send(message, topic, producer)
}

func initProducer() (sarama.SyncProducer, error) {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Version = sarama.V0_11_0_0
	config.ClientID = "go-cleanarch"

	prd, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_HOST")}, config)

	return prd, err
}

func send(message string, topic string, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}
