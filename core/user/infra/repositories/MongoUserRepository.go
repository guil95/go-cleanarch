package user

import (
	"context"
	"fmt"
	userDomain "github.com/guil95/go-cleanarch/core/user/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoUserRepository struct {
	db *mongo.Database
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{
		db: db,
	}
}

func (repo MongoUserRepository) List() (error, *[]userDomain.User) {
	var users []userDomain.User

	ctx := context.Background()

	collection := repo.db.Collection("users")

	cursor, err := collection.Find(ctx, bson.D{{}})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return userDomain.UserNotFound, nil
		}

		return err, nil
	}

	for cursor.Next(ctx) {
		var u userDomain.User

		err := cursor.Decode(&u)
		if err != nil {
			return err, nil
		}

		users = append(users, u)
	}

	err = cursor.Close(ctx)

	if err != nil {
		return err, nil
	}

	return nil, &users
}

func (repo MongoUserRepository) Get(uuid userDomain.UUID) (error, *userDomain.User) {
	var u userDomain.User

	ctx := context.Background()

	collection := repo.db.Collection("users")

	err := collection.FindOne(ctx, bson.D{{Key: "identifier", Value: uuid.String()}}).Decode(&u)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return userDomain.UserNotFound, nil
		}
		return err, nil
	}

	return nil, &u
}

func (repo MongoUserRepository) Create(user *userDomain.User) (error, *userDomain.User) {
	ctx := context.Background()
	fmt.Println(user)
	collection := repo.db.Collection("users")

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err, nil
	}

	return nil, user
}

func (repo MongoUserRepository) CreateBatch(users []*userDomain.User) error {
	ctx := context.Background()

	collection := repo.db.Collection("users")

	var usersInterface []interface{}

	for _, item := range users {
		usersInterface = append(usersInterface, item)
	}

	_, err := collection.InsertMany(ctx, usersInterface)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (repo MongoUserRepository) SearchByName(userName string) (error, *userDomain.User) {
	var u userDomain.User

	ctx := context.Background()

	collection := repo.db.Collection("users")
	err := collection.FindOne(ctx, bson.M{"name": userName}, options.FindOne().SetProjection(bson.M{"_id": 0})).Decode(&u)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return userDomain.UserNotFound, nil
		}
		return err, nil
	}

	return nil, &u
}
