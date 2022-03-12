package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/renanrs/lms-hakka/src/internal/course"
)

func main() {
	lambda.Start(course.CreateHandler)
}
