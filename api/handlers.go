package api

import (
	"context"

	"github.com/Lemm8/AlumnoInfo-CollegeAPI.git/helpers"
	"github.com/aws/aws-lambda-go/events"
)

func GetAlumnosInfo(ctx context.Context, req events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {
	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		return helpers.GetAlumnoInfoResponse()
	}
	return helpers.GetAlumnosInfoResponse()
}

func PostAlumnoInfo(ctx context.Context, req events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {
	return helpers.PostAlumnoInfoResponse()
}

func PutAlumnoInfo(ctx context.Context, req events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {
	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		return helpers.ServerError(400, "Se debe incluir el ID")
	}
	return helpers.PutAlumnoInfoResponse()
}

func DeleteAlumnoInfo(ctx context.Context, req events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {
	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		return helpers.ServerError(400, "Se debe incluir el ID")
	}
	return helpers.DeleteAlumnoInfoResponse()
}
