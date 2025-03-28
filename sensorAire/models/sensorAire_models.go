package models

type SensorData struct {
	Valor float64 `json:"valor"`
}

type SensorDataProcesado struct {
	Valor     float64 `json:"valor"`
	Categoria string  `json:"categoria"`
	Timestamp string  `json:"timestamp"`
}
