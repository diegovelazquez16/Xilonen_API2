package models

type CalidadAire struct {
	Valor float64 `json:"valor"`
}

type CalidadAireProcesado struct {
	Valor     float64 `json:"valor"`
	Categoria string  `json:"categoria"`
	Timestamp string  `json:"timestamp"`
}
 