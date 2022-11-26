package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/Lemm8/AlumnoInfo-CollegeAPI.git/database"
	"github.com/Lemm8/AlumnoInfo-CollegeAPI.git/helpers"
	"github.com/aws/aws-lambda-go/events"
)

func GetAlumnosInfo(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		alumnoInfo := &helpers.AlumnoInfo{}
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, errors.New("el id no es válido")
		}

		alumnoInfo, err = database.FetchAlumnoInfo(ctx, db, intID)
		return helpers.GetAlumnoInfoResponse(alumnoInfo), nil
	}

	alumnosInfo, err := database.FetchAlumnosInfo(ctx, db, req)
	if err != nil {
		return nil, err
	}

	return helpers.GetAlumnosInfoResponse(alumnosInfo), nil

}

func PostAlumnoInfo(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	alumnoInfo := &helpers.AlumnoInfo{}
	err := json.Unmarshal([]byte(req.Body), alumnoInfo)
	if err != nil {
		return helpers.ServerError(500, string(err.Error())), nil
	}

	alumnoInfo, err = database.CreateAlumnoInfo(ctx, db, alumnoInfo)
	if err != nil {
		return helpers.ServerError(500, string(err.Error())), nil
	}

	return helpers.PostAlumnoInfoResponse(alumnoInfo), nil
}

func PutAlumnoInfo(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	alumnoInfo := &helpers.AlumnoInfo{}
	err := json.Unmarshal([]byte(req.Body), &alumnoInfo)
	if err != nil {
		return helpers.ServerError(500, string(err.Error())), nil
	}

	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, errors.New("el id no es válido")
		}
		UpdateAlumnoInfo, err := database.UpdateAlumnoInfo(ctx, db, intID, alumnoInfo)
		if err != nil {
			return nil, err
		}
		return helpers.PutAlumnoInfoResponse(UpdateAlumnoInfo), nil
	}

	return nil, errors.New("se debe incluir el id del alumno_info para actualizarlo")
	// // VALIDATE ID EXISTS IN QUERY
	// id := req.QueryStringParameters["id"]
	// if len(id) < 1 {
	// 	return helpers.ServerError(400, "Se debe incluir el ID"), nil
	// }

	// // CONVERT ID TO INT
	// intID, err := strconv.Atoi(id)
	// if err != nil {
	// 	return nil, errors.New("el id no es válido")
	// }

	// // UNMARSHAL BODY TO ALUMNOINFO
	// alumnoInfo := &helpers.AlumnoInfo{}
	// err = json.Unmarshal([]byte(req.Body), alumnoInfo)
	// if err != nil {
	// 	return helpers.ServerError(500, string(err.Error())), nil
	// }

	// // UPDATE ALUMNO
	// alumnoInfo, err = database.UpdateAlumnoInfo(ctx, db, alumnoInfo, intID)
	// return helpers.PutAlumnoInfoResponse(alumnoInfo), nil
}

func DeleteAlumnoInfo(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// VALIDATE ID EXISTS IN QUERY
	id := req.QueryStringParameters["id"]
	if len(id) < 1 {
		return helpers.ServerError(400, "Se debe incluir el ID"), nil
	}

	// CONVERT ID TO INT
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("el id no es válido")
	}

	// DELETE ALUMNO
	alumnoInfo := &helpers.AlumnoInfo{}

	alumnoInfo, err = database.DeleteAlumnoInfo(ctx, db, intID)
	if err != nil {
		return helpers.ServerError(500, string(err.Error())), nil
	}

	return helpers.DeleteAlumnoInfoResponse(alumnoInfo), nil
}
