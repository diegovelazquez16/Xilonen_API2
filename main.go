package main

import (
	"log"
	aire "Xilonen-2/sensorAire/messaging"
	humedad "Xilonen-2/sensorHumedad/messaging"
)

func main() {
	log.Println("🚀 Iniciando consumidor de datos...")
	go aire.StartConsumer()
	go humedad.StartHumedadConsumer()

	select {}
}
