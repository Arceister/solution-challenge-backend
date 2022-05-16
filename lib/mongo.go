package lib

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Mongo *mongo.Database
	Ctx   *context.Context
}

func NewMongoDatabase(env Env) Mongo {
	MongoURI := env.MongoURI

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(MongoURI).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Mongo Connection Error")
	} else {
		fmt.Println("Mongo Connection Success")
	}

	mongo := client.Database("Kinder-Dev")

	return Mongo{
		Mongo: mongo,
		Ctx:   &ctx,
	}
}
