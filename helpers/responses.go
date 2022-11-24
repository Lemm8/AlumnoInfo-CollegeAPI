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

func GetAlumnosInfoResponse(alumnosInfo []*AlumnoInfo) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&AlumnosInfoStruct{
		Status:      http.StatusOK,
		AlumnosInfo: alumnosInfo,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func GetAlumnoInfoResponse(alumnoInfo *AlumnoInfo) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&AlumnoInfoStruct{
		Status:     http.StatusOK,
		Message:    "OK",
		AlumnoInfo: alumnoInfo,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func PostAlumnoInfoResponse(alumnoInfo *AlumnoInfo) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&AlumnoInfoStruct{
		Status:     http.StatusOK,
		Message:    "AlumnoInfo Created",
		AlumnoInfo: alumnoInfo,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func PutAlumnoInfoResponse(alumnoInfo *AlumnoInfo) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&AlumnoInfoStruct{
		Status:     http.StatusOK,
		Message:    "AlumnoInfo Updated",
		AlumnoInfo: alumnoInfo,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func DeleteAlumnoInfoResponse(alumnoInfo *AlumnoInfo) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&AlumnoInfoStruct{
		Status:     http.StatusOK,
		Message:    "AlumnoIndo Deleted",
		AlumnoInfo: alumnoInfo,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}
