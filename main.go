package main

import (
	"log"
	"Xilonen-2/sensorAire/messaging"
)

func main() {
	log.Println("🚀 Iniciando consumidor de datos...")
	messaging.StartConsumer()
}
