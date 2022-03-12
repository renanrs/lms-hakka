package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri string = "mongodb://root:example@mongo:27017"
)

type Database struct {
	client *mongo.Client
}

func CreateDbInstance() *Database {
	database := &Database{}
	database.initClient()
	return database
}

func (db *Database) initClient() *mongo.Client {
	if db.client != nil {
		return db.client
	}

	var err error
	clientOptions := options.Client().ApplyURI(uri)
	db.client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = db.client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB.")
	return db.client
}

func (db *Database) GetCollection(name string) *mongo.Collection {
	collection := db.client.Database("lms-hakka").Collection(name)
	return collection
}

func (db *Database) CloseConnection() {
	if db.client == nil {
		return
	}

	err := db.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
