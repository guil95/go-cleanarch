package user_consumers

import (
	"context"
	"encoding/json"
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

type CreateUserConsumerGroup struct {
	repo     user.Repository
	consumer sarama.ConsumerGroup
}

func NewCreateUserConsumer(repo user.Repository, consumer sarama.Consumer) *CreateUserConsumer {

	return &CreateUserConsumer{
		repo:     repo,
		consumer: consumer,
	}
}

func NewCreateUserConsumerGroup(repo user.Repository, consumer sarama.ConsumerGroup) *CreateUserConsumerGroup {
	return &CreateUserConsumerGroup{
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
			consumer.HighWaterMarkOffset()
			var users []*user.User
			usersJson := string(msg.Value)

			err := json.Unmarshal([]byte(usersJson), &users)

			if err != nil {
				return err
			}

			_ = c.repo.CreateBatch(users)

			fmt.Println(fmt.Sprintf("%d novos usuarios", len(users)))
		case <-signals:
			break ConsumerLoop
		}
	}

	return nil
}

func (c *CreateUserConsumerGroup) ListenGroup() error {
	ctx := context.Background()
	for {
		topics := []string{topic}
		handler := c

		err := c.consumer.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (c *CreateUserConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (c *CreateUserConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (c *CreateUserConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var users []*user.User
		usersJson := string(msg.Value)

		err := json.Unmarshal([]byte(usersJson), &users)

		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("%d novos usuarios", len(users)))
		go func(users []*user.User) {
			_ = c.repo.CreateBatch(users)
		}(users)

		sess.MarkMessage(msg, "")
	}

	return nil
}
