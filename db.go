package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func db() *mongo.Client {
	// Get the value of an Environment Variable
 	host := os.Getenv("DB_URL")
	log.Println(host)
	clientOptions := options.Client().ApplyURI(host)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
