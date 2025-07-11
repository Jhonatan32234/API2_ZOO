package models


import (
	"api2/db"
	"api2/src/entities"
)


type OjivaResultVisitas struct {
	Hora  string `json:"hora"`
	Total int    `json:"total"`
}

func GetNowVisitas(zona string) ([]entities.Visitas, error) {
	var visitas []entities.Visitas
	err := db.DB.Raw(`
		SELECT * FROM visitas 
		WHERE fecha = (SELECT MAX(fecha) FROM visitas WHERE zona = ?) 
		AND zona = ? 
		AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16
	`, zona, zona).Scan(&visitas).Error
	return visitas, err
}


func GetLastWeekVisitas(zona string) ([]entities.Visitas, error) {
	var fechas []string
	err := db.DB.Raw(`
		SELECT DISTINCT fecha 
		FROM visitas 
		WHERE zona = ?
		ORDER BY fecha DESC 
		LIMIT 6
	`, zona).Scan(&fechas).Error
	if err != nil || len(fechas) == 0 {
		return nil, err
	}

	var visitas []entities.Visitas
	err = db.DB.Where("fecha IN ? AND zona = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16", fechas, zona).Find(&visitas).Error
	return visitas, err
}


func GetYesterdayVisitas(zona string) ([]entities.Visitas, error) {
	var visitas []entities.Visitas
	var fecha string

	err := db.DB.Raw(`
		SELECT DISTINCT fecha 
		FROM visitas 
		WHERE zona = ?
		ORDER BY fecha DESC 
		LIMIT 1 OFFSET 1
	`, zona).Scan(&fecha).Error
	if err != nil || fecha == "" {
		return visitas, err
	}

	err = db.DB.Where("fecha = ? AND zona = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16", fecha, zona).Find(&visitas).Error
	return visitas, err
}


func GetOjivaVisitas(fecha, zona string) ([]OjivaResultVisitas, error) {
	var result []OjivaResultVisitas

	if fecha == "" {
		err := db.DB.Raw(`SELECT MAX(fecha) FROM visitas WHERE zona = ?`, zona).Scan(&fecha).Error
		if err != nil {
			return result, err
		}
	}

	err := db.DB.Raw(`
		SELECT hora, SUM(visitantes) as total 
		FROM visitas 
		WHERE fecha = ? AND zona = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16
		GROUP BY hora 
		ORDER BY hora
	`, fecha, zona).Scan(&result).Error

	return result, err
}


func GetFechaVisitas(fecha, zona string) ([]entities.Visitas, error) {
	var visitas []entities.Visitas
	err := db.DB.Where("fecha = ? AND zona = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16", fecha, zona).Find(&visitas).Error
	return visitas, err
}
