package main

import (
	"log"
	aire "Xilonen-2/sensorAire/messaging"
	humedad "Xilonen-2/sensorHumedad/messaging"
	nivelAgua "Xilonen-2/sensorNivelAgua/messaging"
	sensorUV "Xilonen-2/sensorUV/messaging"
	sensorTemperatura "Xilonen-2/sensorTemperatura/messaging"
)

func main() {
	log.Println("🚀 Iniciando consumidor de datos...")
	go aire.StartConsumer()
	go humedad.StartHumedadConsumer()
	go nivelAgua.StartNivelAguaConsumer()
	go sensorUV.StartUVConsumer()
	go sensorTemperatura.StartTemperaturaConsumer()

	select {}
}
