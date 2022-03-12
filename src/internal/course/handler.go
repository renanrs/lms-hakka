package course

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/renanrs/lms-hakka/src/pkg/response"
)

type CreateCourseInput struct {
	ProductID   uuid.UUID `json:"productId"`
	UserID      uuid.UUID `json:"userId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func CreateHandler(ctx context.Context, e events.APIGatewayProxyRequest) (response.Response, error) {
	course := &CreateCourseInput{}
	err := json.Unmarshal([]byte(e.Body), course)
	if err != nil {
		return response.Error(err)
	}

	res, err := CreateCourse(course)
	if err != nil {
		return response.Error(err)
	}

	return response.Created(res)
}

func GetHandler(ctx context.Context, e events.APIGatewayProxyRequest) (response.Response, error) {
	value, found := e.PathParameters["id"]

	if !found {
		return response.BadRequest("Id is required")
	}

	res, err := GetCourseByID(value)
	if err != nil {
		return response.Error(err)
	}

	return response.Ok(res)
}
