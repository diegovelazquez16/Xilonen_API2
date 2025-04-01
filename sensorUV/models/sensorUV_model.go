package models

type LuzUV struct {
	Valor float64 `json:"valor"`
}

type LuzUVProcesada struct {
	Valor     float64 `json:"valor"`
	Categoria string  `json:"categoria"`
	Timestamp string  `json:"timestamp"`
}
 