package workers

import (
	"fmt"
	"github.com/Shopify/sarama"
	user_consumers "github.com/guil95/go-cleanarch/core/user/infra/brokers/kafka/consumers"
	user "github.com/guil95/go-cleanarch/core/user/infra/repositories"
	"gorm.io/gorm"
	"log"
	"os"
)

func Run(db *gorm.DB) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.ClientID = "go-cleanarch"
	master, err := sarama.NewConsumer([]string{os.Getenv("KAFKA_HOST")}, config)

	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			log.Panic(err)
		}
	}()

	consumersUser(db, master)
}

func consumersUser(db *gorm.DB, consumer sarama.Consumer) {
	fmt.Println("init consumer")
	repo := user.NewMysqlUserRepository(db)

	createUserConsumer := user_consumers.NewCreateUserConsumer(repo, consumer)
	err := createUserConsumer.Listen()

	if err != nil {
		return
	}
}
