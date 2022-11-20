package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func PathDoesNotExist() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusNotFound,
		Message: "This path does not exists",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       string(body),
	}
}

func UnhandledMethod() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusBadRequest,
		Message: "Unhandled method, try again",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       string(body),
	}
}

func ServerError(statuscode int, errMessage string) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  statuscode,
		Message: errMessage,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: statuscode,
		Body:       string(body),
	}
}

func GetAlumnosInfoResponse() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusOK,
		Message: "Get Alumnos Info",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func GetAlumnoInfoResponse() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusOK,
		Message: "Get Alumno Info",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func PostAlumnoInfoResponse() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusOK,
		Message: "Post Alumno Info",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func PutAlumnoInfoResponse() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusOK,
		Message: "Put Alumno Info",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func DeleteAlumnoInfoResponse() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusOK,
		Message: "Delete Alumno Info",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}
