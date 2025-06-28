package models

import (
	"api2/db"
	"api2/src/entities"
)

func GetNowAtraccion() ([]entities.Atraccion, error) {
	var atracciones []entities.Atraccion
	err := db.DB.Raw("SELECT * FROM atraccion WHERE fecha = (SELECT MAX(fecha) FROM atraccion)").Scan(&atracciones).Error
	return atracciones, err
}

func GetLastWeekAtraccion() ([]entities.Atraccion, error) {
	var fechas []string

	err := db.DB.Raw(`
		SELECT DISTINCT fecha 
		FROM atraccion 
		ORDER BY fecha DESC 
		LIMIT 6
	`).Scan(&fechas).Error
	if err != nil || len(fechas) == 0 {
		return nil, err
	}

	var atracciones []entities.Atraccion

	err = db.DB.Where("fecha IN ?", fechas).Find(&atracciones).Error
	return atracciones, err
}


func GetYesterdayAtraccion() ([]entities.Atraccion, error) {
	var atracciones []entities.Atraccion
	var fecha string

	err := db.DB.Raw(`
		SELECT DISTINCT fecha 
		FROM atraccion 
		ORDER BY fecha DESC 
		LIMIT 1 OFFSET 1
	`).Scan(&fecha).Error
	if err != nil || fecha == "" {
		return atracciones, err
	}

	err = db.DB.Where("fecha = ?", fecha).Find(&atracciones).Error
	return atracciones, err
}


type OjivaResultAtraccion struct {
	Hora  string `json:"hora"`
	Total int    `json:"total"`
}

func GetOjivaAtraccion(fecha string) ([]OjivaResultAtraccion, error) {
	var result []OjivaResultAtraccion

	if fecha == "" {
		err := db.DB.Raw(`
			SELECT MAX(fecha) 
			FROM atraccion
		`).Scan(&fecha).Error
		if err != nil {
			return result, err
		}
	}

	err := db.DB.Raw(`
		SELECT hora, SUM(tiempo) as total 
		FROM atraccion 
		WHERE fecha = ? 
		GROUP BY hora 
		ORDER BY hora
	`, fecha).Scan(&result).Error

	return result, err
}

func GetFechaAtraccion(fecha string) ([]entities.Atraccion, error) {
	var atracciones []entities.Atraccion
	err := db.DB.Where("fecha = ?", fecha).Find(&atracciones).Error
	return atracciones, err
}
