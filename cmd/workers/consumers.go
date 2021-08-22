package workers

import (
	"fmt"
	"github.com/Shopify/sarama"
	user_consumers "github.com/guil95/go-cleanarch/core/user/infra/brokers/kafka/consumers"
	user "github.com/guil95/go-cleanarch/core/user/infra/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func Run(db *mongo.Database) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.ClientID = "go-cleanarch"
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err := sarama.NewConsumerGroup([]string{os.Getenv("KAFKA_HOST")}, "go-cleanarch", config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	consumeUserGroup(db, group)
}

func consumersUser(db *mongo.Database, consumer sarama.Consumer) {
	fmt.Println("init consumer")
	repo := user.NewMongoUserRepository(db)

	createUserConsumer := user_consumers.NewCreateUserConsumer(repo, consumer)
	err := createUserConsumer.Listen()

	if err != nil {
		return
	}
}

func consumeUserGroup(db *mongo.Database, consumer sarama.ConsumerGroup) {
	fmt.Println("init consumer")
	repo := user.NewMongoUserRepository(db)

	createUserConsumer := user_consumers.NewCreateUserConsumerGroup(repo, consumer)
	err := createUserConsumer.ListenGroup()

	if err != nil {
		return
	}
}
