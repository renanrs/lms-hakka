package course

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCourse(input *CreateCourseInput) (*InsertOneResult, error) {
	course := &Course{
		UserID:      input.UserID.String(),
		ProductID:   input.ProductID.String(),
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
	}

	res, err := InsertOne(course)
	return res, err
}

func GetCourseByID(id string) (*Course, error) {

	docId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": docId}
	res, err := FindOneByFilter(filter)

	return res, err
}
