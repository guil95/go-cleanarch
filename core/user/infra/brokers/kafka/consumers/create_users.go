package user_consumers

import (
	"fmt"
	"github.com/Shopify/sarama"
	user "github.com/guil95/go-cleanarch/core/user/domain"
	"log"
	"os"
	"os/signal"
)

const topic = "users_create"

type CreateUserConsumer struct {
	repo     user.Repository
	consumer sarama.Consumer
}

func NewCreateUserConsumer(repo user.Repository, consumer sarama.Consumer) *CreateUserConsumer {

	return &CreateUserConsumer{
		repo:     repo,
		consumer: consumer,
	}
}

func (c *CreateUserConsumer) Listen() error {
	consumer, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

ConsumerLoop:
	for {
		select {
		case msg := <-consumer.Messages():
			fmt.Println(msg.Value)
		case <-signals:
			break ConsumerLoop
		}
	}

	return nil
}
