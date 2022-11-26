package validators

import (
	"context"
	"database/sql"

	"github.com/Lemm8/AlumnoInfo-CollegeAPI.git/helpers"
)

const getAlumnoQuery = `SELECT * FROM Alumno WHERE ID = ?;`
const getMateriaQuery = `SELECT * FROM Materia WHERE ID = ?;`
const getCalificacionQuery = `SELECT * FROM Calificacion WHERE ID = ?;`

func IsValidPath(path string) bool {

	if path != "/AlumnoInfo" {
		return false
	}

	return true
}

func AlumnoExists(ctx context.Context, db *sql.DB, id int) bool {
	row := db.QueryRowContext(ctx, getAlumnoQuery, id)

	alumno := &helpers.Alumno{}
	if err := row.Scan(&alumno.ID, &alumno.Nombre, &alumno.Apellido, &alumno.Matricula,
		&alumno.Fecha_Nacimiento, &alumno.Semestre, &alumno.Carreras_ID); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return true
}

func CalificacionExists(ctx context.Context, db *sql.DB, id int) bool {
	// QUERY CALIFICACION BY ID
	row := db.QueryRowContext(ctx, getCalificacionQuery, id)

	calificacion := &helpers.Calificacion{}
	if err := row.Scan(&calificacion.ID, &calificacion.Evaluacion); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return true
}

func MateriaExists(ctx context.Context, db *sql.DB, id int) bool {
	// QUERY CALIFICACION BY ID
	row := db.QueryRowContext(ctx, getMateriaQuery, id)

	materia := &helpers.Materia{}
	if err := row.Scan(&materia.ID, &materia.Nombre, &materia.IsTroncoComun); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return true
}
