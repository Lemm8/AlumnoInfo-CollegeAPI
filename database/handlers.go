package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Lemm8/AlumnoInfo-CollegeAPI.git/helpers"
	"github.com/Lemm8/AlumnoInfo-CollegeAPI.git/validators"
	"github.com/aws/aws-lambda-go/events"
)

const getAlumnosInfoQuery = `SELECT * FROM AlumnoInfo;`

const getAlumnoInfoQuery = `SELECT * FROM AlumnoInfo WHERE ID = ?;`

const insertAlumnoInfoSQL = `INSERT INTO AlumnoInfo (Estado, Materia_ID, Calificacion_ID, Alumno_ID) 
VALUES (?, ?, ?, ?);`

const updateAlumnoInfoSQL = `UPDATE AlumnoInfo SET Estado = ?, Materia_ID = ?, Calificacion_ID = ?, Alumno_ID = ? WHERE ID = ?;`

const deleteAlumnoInfoSQL = `DELETE FROM AlumnoInfo WHERE ID = ?;`

func FetchAlumnosInfo(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) ([]*helpers.AlumnoInfo, error) {
	rows, err := db.QueryContext(ctx, getAlumnoInfoQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	alumnosInfo := make([]*helpers.AlumnoInfo, 0)

	for rows.Next() {
		alumnoInfo := &helpers.AlumnoInfo{}
		if err := rows.Scan(&alumnoInfo.ID, &alumnoInfo.Estado, &alumnoInfo.Calificacion_ID); err != nil {
			return nil, err
		}
		alumnosInfo = append(alumnosInfo, alumnoInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return alumnosInfo, nil
}

func FetchAlumnoInfo(ctx context.Context, db *sql.DB, id int) (*helpers.AlumnoInfo, error) {
	row := db.QueryRowContext(ctx, getAlumnoInfoQuery, id)

	alumnoInfo := &helpers.AlumnoInfo{}
	if err := row.Scan(&alumnoInfo.ID, &alumnoInfo.Estado, &alumnoInfo.Calificacion_ID,
		&alumnoInfo.Materia_ID, &alumnoInfo.Alumno_ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("No existe un docente con el ID %v", id))
		}
		return nil, err
	}
	return alumnoInfo, nil
}

func CreateAlumnoInfo(ctx context.Context, db *sql.DB, alumnoInfo *helpers.AlumnoInfo) (*helpers.AlumnoInfo, error) {
	if !validators.AlumnoExists(ctx, db, alumnoInfo.Alumno_ID) {
		return nil, errors.New(fmt.Sprintf("No existe un alumno con el ID %v", alumnoInfo.Alumno_ID))
	}
	if !validators.MateriaExists(ctx, db, alumnoInfo.Materia_ID) {
		return nil, errors.New(fmt.Sprintf("No existe una materia con el ID %v", alumnoInfo.Materia_ID))
	}
	if !validators.CalificacionExists(ctx, db, alumnoInfo.Calificacion_ID) {
		return nil, errors.New(fmt.Sprintf("No existe una calificación con el ID %v", alumnoInfo.Calificacion_ID))
	}

	res, err := db.ExecContext(ctx, insertAlumnoInfoSQL, alumnoInfo.Estado, alumnoInfo.Materia_ID,
		alumnoInfo.Calificacion_ID, alumnoInfo.Alumno_ID)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	alumnoInfo = &helpers.AlumnoInfo{
		ID:              int(id),
		Estado:          alumnoInfo.Estado,
		Materia_ID:      alumnoInfo.Materia_ID,
		Calificacion_ID: alumnoInfo.Calificacion_ID,
		Alumno_ID:       alumnoInfo.Alumno_ID,
	}

	return alumnoInfo, nil
}

func UpdateAlumnoInfo(ctx context.Context, db *sql.DB, alumnoInfo *helpers.AlumnoInfo, id int) (*helpers.AlumnoInfo, error) {

	_, err := FetchAlumnoInfo(ctx, db, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("No existe un registro con el ID: %v", id))
	}

	if !validators.AlumnoExists(ctx, db, alumnoInfo.Alumno_ID) {
		return nil, errors.New(fmt.Sprintf("No existe un alumno con el ID: %v", alumnoInfo.Alumno_ID))
	}
	if !validators.CalificacionExists(ctx, db, alumnoInfo.Calificacion_ID) {
		return nil, errors.New(fmt.Sprintf("No existe una calificacion con el ID: %v", alumnoInfo.Alumno_ID))
	}
	if !validators.MateriaExists(ctx, db, alumnoInfo.Materia_ID) {
		return nil, errors.New(fmt.Sprintf("No existe una materia con el ID: %v", alumnoInfo.Alumno_ID))
	}

	_, err = db.ExecContext(ctx, updateAlumnoInfoSQL, alumnoInfo.Estado, alumnoInfo.Materia_ID,
		alumnoInfo.Calificacion_ID, alumnoInfo.Alumno_ID, alumnoInfo.ID)

	if err != nil {
		return nil, err
	}

	updatedAlumnoInfo := helpers.AlumnoInfo{
		ID:              alumnoInfo.ID,
		Estado:          alumnoInfo.Estado,
		Calificacion_ID: alumnoInfo.Calificacion_ID,
		Materia_ID:      alumnoInfo.Materia_ID,
		Alumno_ID:       alumnoInfo.Alumno_ID,
	}

	return &updatedAlumnoInfo, nil
}

func DeleteAlumnoInfo(ctx context.Context, db *sql.DB, id int) (*helpers.AlumnoInfo, error) {
	alumnoInfo, err := FetchAlumnoInfo(ctx, db, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("No existe un registro con el ID: %v", id))
	}

	_, err = db.ExecContext(ctx, deleteAlumnoInfoSQL, id)
	if err != nil {
		return nil, err
	}

	return alumnoInfo, nil
}
