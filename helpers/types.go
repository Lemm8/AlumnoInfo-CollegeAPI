package helpers

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type AlumnoInfo struct {
	ID              int    `json:"id"`
	Estado          string `json:"estado"`
	Materia_ID      int    `json:"materia_id"`
	Calificacion_ID int    `json:"calificacion_id"`
	Alumno_ID       int    `json:"alumno_id"`
}

type AlumnosInfoStruct struct {
	Status      int           `json:"status"`
	AlumnosInfo []*AlumnoInfo `json:"alumnosInfo"`
}

type AlumnoInfoStruct struct {
	Status     int         `json:"status"`
	Message    string      `json:"msg"`
	AlumnoInfo *AlumnoInfo `json:"alumnoInfo"`
}

type Alumno struct {
	ID               int    `json:"id"`
	Nombre           string `json:"nombre"`
	Apellido         string `json:"apellido"`
	Matricula        string `json:"matricula"`
	Fecha_Nacimiento string `json:"fecha_nacimiento"`
	Semestre         string `json:"semestre"`
}

type Calificacion struct {
	ID         int `json:"id"`
	Evaluacion int `json:"evaluacion"`
}

type Materia struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	IsTroncoComun bool   `json:"troncoComun"`
}
