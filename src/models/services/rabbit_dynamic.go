package services

import (
	"api2/db"
	"api2/src/entities"
	"api2/utils"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)


var (
	zonaConsumers     = make(map[string]bool)
	zonaConsumersLock = &sync.Mutex{}
)

func StartDynamicConsumerByZona(zona string) {
	zonaConsumersLock.Lock()
	defer zonaConsumersLock.Unlock()

	if zonaConsumers[zona] {
		log.Printf("⚠️ Consumidor ya iniciado para la zona: %s", zona)
		return
	}

	zonaConsumers[zona] = true
	log.Printf("🚀 Iniciando consumidores para la zona: %s", zona)

	go consumeZonaTopic("visitas_topic", fmt.Sprintf("visitas.%s", zona), handleZonaVisita)
	go consumeZonaTopic("atracciones_topic", fmt.Sprintf("atracciones.%s", zona), handleZonaAtraccion)
}


func consumeZonaTopic(exchange, routingKey string, handler func(uint)) {
	log.Printf("📡 Iniciando consumidor para zona: exchange='%s', routingKey='%s'\n", exchange, routingKey)

	conn, err := amqp.Dial("amqp://admin:password@localhost:5672/")
	if err != nil {
		log.Println("❌ RabbitMQ conexión fallida:", err)
		return
	}
	log.Println("✅ Conectado a RabbitMQ para zona.")
	ch, _ := conn.Channel()

	err = ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
	if err != nil {
		log.Println("❌ Error declarando exchange:", err)
		return
	}

	q, err := ch.QueueDeclare("", false, true, true, false, nil)
	if err != nil {
		log.Println("❌ Error declarando cola:", err)
		return
	}

	err = ch.QueueBind(q.Name, routingKey, exchange, false, nil)
	if err != nil {
		log.Println("❌ Error enlazando cola a tópico:", err)
		return
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Println("❌ Error al consumir:", err)
		return
	}

	go func() {
		log.Println("🟢 Esperando mensajes de zona en:", routingKey)
		for d := range msgs {
			log.Printf("📥 Mensaje recibido de zona (%s): %s\n", routingKey, string(d.Body))

			var payload struct {
				Id uint `json:"id"`
			}
			if err := json.Unmarshal(d.Body, &payload); err != nil {
				log.Println("❌ Error parsing ID payload:", err)
				continue
			}

			log.Printf("🔍 ID extraído del mensaje: %d\n", payload.Id)

			go func(id uint) {
				time.Sleep(1 * time.Second)
				handler(id)
			}(payload.Id)
		}
	}()
}


func handleZonaVisita(id uint) {
	log.Printf("📌 Procesando visita por zona con ID: %d\n", id)

	var v entities.Visitas
	result := db.DB.Where("id = ? AND enviado = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16", id, false).First(&v)
	if result.Error != nil {
		log.Printf("❌ No se encontró la visita válida con ID %d: %v\n", id, result.Error)
		return
	}

	log.Printf("📄 Visita encontrada: %+v\n", v)

	log.Println("📤 Enviando visita al WebSocket...")
	var visitaMap map[string]interface{}
	jsonBytes, _ := json.Marshal(v)
	json.Unmarshal(jsonBytes, &visitaMap)

	utils.NotifyClients(map[string]interface{}{
		"type": "visita",
		"data": visitaMap,
	})
	log.Println("✅ Enviado al WebSocket.")

	// ✅ Marcar como enviada en la base de datos
	if err := db.DB.Model(&entities.Visitas{}).Where("id = ?", v.Id).Update("enviado", true).Error; err != nil {
		log.Printf("❌ Error al actualizar campo 'enviado' para visita ID %d: %v\n", v.Id, err)
	} else {
		log.Printf("🟢 Visita ID %d actualizada como enviada.\n", v.Id)
	}
}


func handleZonaAtraccion(id uint) {
	log.Printf("📌 Procesando atracción por zona con ID: %d\n", id)

	var a entities.Atraccion
	result := db.DB.Where("id = ? AND enviado = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16", id, false).First(&a)
	if result.Error != nil {
		log.Printf("❌ No se encontró la atracción válida con ID %d: %v\n", id, result.Error)
		return
	}

	log.Printf("📄 Atracción encontrada: %+v\n", a)

	log.Println("📤 Enviando atracción al WebSocket...")
	var atraccionMap map[string]interface{}
	jsonBytes, _ := json.Marshal(a)
	json.Unmarshal(jsonBytes, &atraccionMap)

	utils.NotifyClients(map[string]interface{}{
		"type": "atraccion",
		"data": atraccionMap,
	})
	log.Println("✅ Enviado al WebSocket.")

	// ✅ Marcar como enviada en la base de datos
	if err := db.DB.Model(&entities.Atraccion{}).Where("id = ?", a.Id).Update("enviado", true).Error; err != nil {
		log.Printf("❌ Error al actualizar campo 'enviado' para atracción ID %d: %v\n", a.Id, err)
	} else {
		log.Printf("🟢 Atracción ID %d actualizada como enviada.\n", a.Id)
	}
}

