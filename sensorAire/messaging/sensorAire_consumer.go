package messaging

import (
	"encoding/json"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"Xilonen-2/core"
	"Xilonen-2/sensorAire/models"
)

// Umbrales de calidad del aire
const (
	UMBRAL_BUENO     = 100.0
	UMBRAL_MODERADO  = 200.0
	UMBRAL_PELIGROSO = 300.0
)

func StartConsumer() {
	core.LoadEnv()

	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		log.Fatal("❌ RABBITMQ_URL no está configurado en las variables de entorno")
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("❌ Error al conectar a RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ Error al abrir un canal en RabbitMQ: %v", err)
	}
	defer ch.Close()

	qProcesado, err := ch.QueueDeclare(
		"aire.procesado", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("❌ Error al declarar la cola aire.procesado: %v", err)
	}

	msgs, err := ch.Consume(
		"sensor.aire", "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("❌ Error al consumir mensajes: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			var sensorData models.SensorData
			if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
				log.Printf("⚠️ Error al deserializar el mensaje: %v", err)
				continue
			}

			categoria := "Bueno"
			if sensorData.Valor > UMBRAL_PELIGROSO {
				categoria = "Peligroso"
			} else if sensorData.Valor > UMBRAL_MODERADO {
				categoria = "Moderado"
			}

			datoProcesado := models.SensorDataProcesado{
				Valor:     sensorData.Valor,
				Categoria: categoria,
				Timestamp: time.Now().Format("2006-01-02 15:04:05"),
			}

			procesadoJSON, err := json.Marshal(datoProcesado)
			if err != nil {
				log.Printf("❌ Error al convertir datos procesados a JSON: %v", err)
				continue
			}

			err = ch.Publish(
				"", qProcesado.Name, false, false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        procesadoJSON,
				})
			if err != nil {
				log.Printf("❌ Error al publicar datos procesados: %v", err)
			} else {
				log.Printf("✅ Dato procesado enviado: Valor=%.2f, Categoría=%s", datoProcesado.Valor, datoProcesado.Categoria)
			}
		}
	}()

	log.Println("📡 Esperando datos del sensor MQ-135...")
	<-forever
}
