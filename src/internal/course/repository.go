package course

import (
	"context"
	"time"

	db "github.com/renanrs/lms-hakka/src/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Course struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
	ProductID   string             `bson:"productId"`
	UserID      string             `bson:"userId"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}

type InsertOneResult mongo.InsertOneResult

var client *db.Database

func getClient() *db.Database {
	if client != nil {
		return client
	}

	client = db.CreateDbInstance()
	return client
}

func disconnect() {
	client.CloseConnection()
}

func getCollection() *mongo.Collection {
	collection := getClient().GetCollection("course")
	return collection
}

func FindOneByFilter(filter bson.M) (*Course, error) {
	defer disconnect()
	var course *Course
	collection := getCollection()
	err := collection.FindOne(context.Background(), filter).Decode(&course)
	return course, err
}

func InsertOne(course *Course) (*InsertOneResult, error) {
	defer disconnect()
	var res *InsertOneResult = &InsertOneResult{}
	collection := getCollection()
	course.ID = primitive.NewObjectID()
	result, err := collection.InsertOne(context.Background(), course)
	if err == nil {
		res.InsertedID = result.InsertedID
	}

	return res, err
}
