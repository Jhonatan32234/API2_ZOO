package services

import (
	"api2/src/entities"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func StartRabbitConsumers() {
	conn, err := amqp.Dial("amqp://admin:password@localhost:5672/")
	if err != nil {
		log.Fatal("RabbitMQ connection failed:", err)
	}
	ch, _ := conn.Channel()

	go consumeVisitas(ch)
	go consumeAtracciones(ch)
}

func consumeVisitas(ch *amqp.Channel) {
	q, _ := ch.QueueDeclare("visitas_queue", true, false, false, false, nil)
	msgs, _ := ch.Consume(q.Name, "", false, false, false, false, nil)

	go func() {
		for d := range msgs {
			log.Println("📨 Mensaje recibido (visitas):", string(d.Body))

			var visitas []entities.Visitas
			if err := json.Unmarshal(d.Body, &visitas); err != nil {
				log.Println("❌ Error al deserializar visitas:", err)
				continue
			}

			guardadas, err := SaveVisitas(visitas)
			if err != nil {
				log.Println("❌ Error al guardar visitas:", err)
				continue
			}

			log.Printf("✅ Visitas guardadas: %d", len(guardadas))
			// Notificar solo las guardadas
			/*for _, visita := range guardadas {
				dbVisita, err := GetVisitaByID(uint(visita.Id))
				if err == nil {
					websocket.NotifyClients(struct {
						Datos *models.Visitas `json:"datos"`
					}{dbVisita})
				}
			}*/

			ch.Ack(d.DeliveryTag, false)
		}
	}()
}

func consumeAtracciones(ch *amqp.Channel) {
	q, _ := ch.QueueDeclare("atracciones_queue", true, false, false, false, nil)
	msgs, _ := ch.Consume(q.Name, "", false, false, false, false, nil)

	go func() {
		for d := range msgs {
			log.Println("📨 Mensaje recibido (atracciones):", string(d.Body))

			var atracciones []entities.Atraccion
			if err := json.Unmarshal(d.Body, &atracciones); err != nil {
				log.Println("❌ Error al deserializar atracciones:", err)
				continue
			}

			guardadas, err := SaveAtracciones(atracciones)
			if err != nil {
				log.Println("❌ Error al guardar atracciones:", err)
				continue
			}
			log.Printf("✅ Atracciones guardadas: %d", len(guardadas))
			/*for _, atr := range guardadas {
				dbAtr, err := GetAtraccionByID(uint(atr.Id))
				if err == nil {
					websocket.NotifyClients(struct {
						Tipo  string             `json:"tipo"`
						Datos *models.Atraccion  `json:"datos"`
					}{"atraccion", dbAtr})
				}
			}*/

			ch.Ack(d.DeliveryTag, false)
		}
	}()
}
