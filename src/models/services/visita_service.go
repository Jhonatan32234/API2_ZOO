package services

import (
	"api2/db"
	"api2/src/entities"
	"log"
	"api2/websocket"
)

func SaveVisitas(input []entities.Visitas) ([]entities.Visitas, error) {
	var guardadas []entities.Visitas

	for _, item := range input {
		item.Enviado = false
		if err := db.DB.Create(&item).Error; err != nil {
			log.Println("❌ Error al guardar visita:", err)
		} else {
			guardadas = append(guardadas, item)

			go func(id uint) {
				if visita, err := GetVisitaByID(id); err == nil {
					websocket.NotifyClients(map[string]interface{}{
						"type": "visita",
						"data": visita,
					})

					db.DB.Model(&entities.Visitas{}).Where("id = ?", id).Update("enviado", true)
				}
			}(uint(item.Id))
		}
	}

	if len(guardadas) == 0 {
		log.Println("⚠️ Ninguna visita fue guardada.")
		return nil, nil
	}

	return guardadas, nil
}


func GetVisitaByID(id uint) (*entities.Visitas, error) {
	var visita entities.Visitas
	if err := db.DB.First(&visita, id).Error; err != nil {
		log.Println("❌ Error al obtener visita por ID:", err)
		return nil, err
	}
	return &visita, nil
}
