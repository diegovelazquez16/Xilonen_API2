package models

type Humedad struct {
	Valor float64 `json:"valor"`
}

type HumedadProcesada struct {
	Valor     float64 `json:"valor"`
	Categoria string  `json:"categoria"`
	Timestamp string  `json:"timestamp"`
}
 