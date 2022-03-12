package response

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
)

type Response events.APIGatewayProxyResponse

func response(model interface{}, statusCode int) (Response, error) {
	responseEmpty := Response{}
	b, err := json.MarshalIndent(model, "", " ")
	if err != nil {
		return responseEmpty, errors.WithStack(err)
	}

	return Response{
		Body: string(b) + "\n",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: statusCode,
	}, nil
}

func Ok(model interface{}) (Response, error) {
	return response(model, 200)
}

func Created(model interface{}) (Response, error) {
	return response(model, 201)
}

func NotFound(model interface{}) (Response, error) {
	return response(model, 404)
}

func BadRequest(model interface{}) (Response, error) {
	return response(model, 400)
}

func Error(err error) (Response, error) {
	return Response{}, errors.WithStack(err)
}
