package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://sanket:Sanket123@service.7rseb.mongodb.net/Service?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	client.Connect(ctx)
	return client
}
