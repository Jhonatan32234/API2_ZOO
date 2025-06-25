package models


import (
	"api2/db"
	"api2/src/entities"
)

func GetNowVisitas() ([]entities.Visitas, error) {
	var visitas []entities.Visitas
	err := db.DB.Raw("SELECT * FROM visitas WHERE fecha = (SELECT MAX(fecha) FROM visitas)").Scan(&visitas).Error
	return visitas, err
}

func GetLastWeekVisitas() ([]entities.Visitas, error) {
	var visitas []entities.Visitas
	err := db.DB.Raw(`
		SELECT * FROM visitas 
		WHERE fecha IN (
			SELECT DISTINCT fecha 
			FROM visitas 
			ORDER BY fecha DESC 
			LIMIT 6
		)
	`).Scan(&visitas).Error
	return visitas, err
}


func GetYesterdayVisitas() ([]entities.Visitas, error) {
	var visitas []entities.Visitas
	var fecha string

	// Obtener la segunda fecha más reciente
	err := db.DB.Raw(`
		SELECT DISTINCT fecha 
		FROM visitas 
		ORDER BY fecha DESC 
		LIMIT 1 OFFSET 1
	`).Scan(&fecha).Error
	if err != nil || fecha == "" {
		return visitas, err
	}

	err = db.DB.Where("fecha = ?", fecha).Find(&visitas).Error
	return visitas, err
}

type OjivaResultVisitas struct {
	Hora  string `json:"hora"`
	Total int    `json:"total"`
}

func GetOjivaVisitas(fecha string) ([]OjivaResultVisitas, error) {
	var result []OjivaResultVisitas

	// Si no se pasa una fecha, obtener la fecha más reciente
	if fecha == "" {
		err := db.DB.Raw(`
			SELECT MAX(fecha) 
			FROM visitas
		`).Scan(&fecha).Error
		if err != nil {
			return result, err
		}
	}

	err := db.DB.Raw(`
		SELECT hora, SUM(visitantes) as total 
		FROM visitas 
		WHERE fecha = ? 
		GROUP BY hora 
		ORDER BY hora
	`, fecha).Scan(&result).Error

	return result, err
}


func GetFechaVisitas(fecha string) ([]entities.Visitas, error) {
	var visitas []entities.Visitas
	err := db.DB.Where("fecha = ?", fecha).Find(&visitas).Error
	return visitas, err
}