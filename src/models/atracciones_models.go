package models

import (
	"api2/db"
	"api2/src/entities"
)


type OjivaResultAtraccion struct {
	Hora  string `json:"hora"`
	Total int    `json:"total"`
}


func GetNowAtraccion(zona string) ([]entities.Atraccion, error) {
	var atracciones []entities.Atraccion
	err := db.DB.Raw(`
		SELECT * FROM atraccion 
		WHERE fecha = (SELECT MAX(fecha) FROM atraccion WHERE zona = ?) 
		AND zona = ?
		AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16
	`, zona, zona).Scan(&atracciones).Error
	return atracciones, err
}


func GetLastWeekAtraccion(zona string) ([]entities.Atraccion, error) {
	var fechas []string
	err := db.DB.Raw(`
		SELECT DISTINCT fecha 
		FROM atraccion 
		WHERE zona = ?
		ORDER BY fecha DESC 
		LIMIT 6
	`, zona).Scan(&fechas).Error
	if err != nil || len(fechas) == 0 {
		return nil, err
	}

	var atracciones []entities.Atraccion
	err = db.DB.Where("fecha IN ? AND zona = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16", fechas, zona).Find(&atracciones).Error
	return atracciones, err
}


func GetYesterdayAtraccion(zona string) ([]entities.Atraccion, error) {
	var atracciones []entities.Atraccion
	var fecha string

	err := db.DB.Raw(`
		SELECT DISTINCT fecha 
		FROM atraccion 
		WHERE zona = ?
		ORDER BY fecha DESC 
		LIMIT 1 OFFSET 1
	`, zona).Scan(&fecha).Error
	if err != nil || fecha == "" {
		return atracciones, err
	}

	err = db.DB.Where("fecha = ? AND zona = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16", fecha, zona).Find(&atracciones).Error
	return atracciones, err
}


func GetOjivaAtraccion(fecha, zona string) ([]OjivaResultAtraccion, error) {
	var result []OjivaResultAtraccion

	if fecha == "" {
		err := db.DB.Raw(`SELECT MAX(fecha) FROM atraccion WHERE zona = ?`, zona).Scan(&fecha).Error
		if err != nil {
			return result, err
		}
	}

	err := db.DB.Raw(`
		SELECT hora, SUM(tiempo) as total 
		FROM atraccion 
		WHERE fecha = ? AND zona = ? AND 
		CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16
		GROUP BY hora 
		ORDER BY hora
	`, fecha, zona).Scan(&result).Error

	return result, err
}


func GetFechaAtraccion(fecha, zona string) ([]entities.Atraccion, error) {
	var atracciones []entities.Atraccion
	err := db.DB.Where("fecha = ? AND zona = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16", fecha, zona).Find(&atracciones).Error
	return atracciones, err
}
