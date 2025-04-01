package messaging

import (
	"encoding/json"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"Xilonen-2/core"
	"Xilonen-2/sensorHumedad/models"
)

const (
	UMBRAL_BAJA  = 30.0
	UMBRAL_MEDIA = 60.0
	UMBRAL_ALTA  = 80.0
)

func StartHumedadConsumer() {
	core.LoadEnv()
	rabbitURL := os.Getenv("RABBITMQ_URL")

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

	qProcesado, err := ch.QueueDeclare("humedad.procesado", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("❌ Error al declarar la cola humedad.procesado: %v", err)
	}

	msgs, err := ch.Consume("sensor.humedad", "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("❌ Error al consumir mensajes: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			var sensorData models.Humedad
			if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
				log.Printf("⚠️ Error al deserializar el mensaje: %v", err)
				continue
			}

			categoria := "Normal"
			if sensorData.Valor < UMBRAL_BAJA {
				categoria = "Baja"
			} else if sensorData.Valor > UMBRAL_ALTA {
				categoria = "Alta"
			}

			datoProcesado := models.HumedadProcesada{
				Valor:     sensorData.Valor,
				Categoria: categoria,
				Timestamp: time.Now().Format("2006-01-02 15:04:05"),
			}

			procesadoJSON, err := json.Marshal(datoProcesado)
			if err != nil {
				log.Printf("❌ Error al convertir datos procesados a JSON: %v", err)
				continue
			}

			err = ch.Publish("", qProcesado.Name, false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        procesadoJSON,
			})
			if err != nil {
				log.Printf("❌ Error al publicar datos procesados: %v", err)
			} else {
				log.Printf("✅ Dato procesado (Humedad) enviado: Valor=%.2f, Categoría=%s", datoProcesado.Valor, datoProcesado.Categoria)
			}
		}
	}()

	log.Println("📡 Esperando datos del sensor de humedad...")
	<-forever
}
