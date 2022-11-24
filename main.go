package main

import (
	"context"

	// "github.com/Lemm8/AlumnoInfo-CollegeAPI.git/database"

	"github.com/Lemm8/AlumnoInfo-CollegeAPI.git/api"
	"github.com/Lemm8/AlumnoInfo-CollegeAPI.git/database"
	"github.com/Lemm8/AlumnoInfo-CollegeAPI.git/helpers"
	"github.com/Lemm8/AlumnoInfo-CollegeAPI.git/validators"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	if !validators.IsValidPath(event.Path) {
		return helpers.PathDoesNotExist(), nil
	}

	// CONNECT TO DB
	dbConnection, err := database.GetConnection()
	_, err = database.GetConnection()
	if err != nil {
		return helpers.ServerError(500, string(err.Error())), nil
	}
	db := dbConnection

	switch event.HTTPMethod {
	case "GET":
		apiRespone, err := api.GetAlumnosInfo(ctx, db, event)
		if err != nil {
			return helpers.ServerError(500, string(err.Error())), nil
		}
		return apiRespone, nil
	case "POST":
		apiRespone, err := api.GetAlumnosInfo(ctx, db, event)
		if err != nil {
			return helpers.ServerError(500, string(err.Error())), nil
		}
		return apiRespone, nil
	case "PUT":
		apiRespone, err := api.GetAlumnosInfo(ctx, db, event)
		if err != nil {
			return helpers.ServerError(500, string(err.Error())), nil
		}
		return apiRespone, nil
	case "DELETE":
		apiRespone, err := api.GetAlumnosInfo(ctx, db, event)
		if err != nil {
			return helpers.ServerError(500, string(err.Error())), nil
		}
		return apiRespone, nil
	default:
		return helpers.UnhandledMethod(), nil
	}

}

func main() {
	lambda.Start(handler)
}
