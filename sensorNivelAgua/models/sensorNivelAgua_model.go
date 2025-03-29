package models

type NivelAgua struct {
	Valor float64 `json:"valor"`
}

type NivelAguaProcesado struct {
	Valor     float64 `json:"valor"`
	Categoria string  `json:"categoria"`
	Timestamp string  `json:"timestamp"`
}
 