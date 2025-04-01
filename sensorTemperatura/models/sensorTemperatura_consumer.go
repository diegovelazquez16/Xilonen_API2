package models

type Temperatura struct {
	Valor float64 `json:"valor"`
}

type TemperaturaProcesada struct {
	Valor     float64 `json:"valor"`
	Categoria string  `json:"categoria"`
	Timestamp string  `json:"timestamp"`
}
 