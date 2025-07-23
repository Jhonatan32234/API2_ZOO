package models


import (
	"api2/db"
)


type NowVisitas struct {
	Fecha string `json:"fecha"`
	Hora  string `json:"hora"`
	Total int    `json:"total"`
}

type LastWeekVisitas struct {
	Fecha string `json:"fecha"`
	Total int    `json:"total"`
}

type YesterdayVisitas struct {
	Fecha string `json:"fecha"`
	Zona  string `json:"zona"`
	Total int    `json:"total"`
}

type OjivaVisitas struct {
	Fecha string `json:"fecha"`
	Hora  string `json:"hora"`
	Total int    `json:"total"`
}

// Acumulado por hora (fecha más reciente)
func GetNowVisitas(zona string) ([]NowVisitas, error) {
	var result []NowVisitas
	err := db.DB.Raw(`
		WITH por_hora AS (
			SELECT 
				fecha,
				CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) AS hora,
				SUM(visitantes) AS total
			FROM visitas
			WHERE fecha = (
				SELECT MAX(fecha) FROM visitas WHERE zona = ?
			)
			AND zona = ?
			AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16
			GROUP BY fecha, hora
		),
		acumulado AS (
			SELECT 
				fecha,
				hora,
				SUM(total) OVER (ORDER BY hora) AS total
			FROM por_hora
		)
		SELECT 
			fecha,
			CONCAT(LPAD(hora, 2, '0'), ':00') AS hora,
			total
		FROM acumulado
		ORDER BY hora
	`, zona, zona).Scan(&result).Error
	return result, err
}

// Total diario de las últimas 6 fechas
func GetLastWeekVisitas(zona string) ([]LastWeekVisitas, error) {
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

	var result []LastWeekVisitas
	err = db.DB.Raw(`
		SELECT fecha, SUM(visitantes) as total
		FROM visitas
		WHERE fecha IN (?) AND zona = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16
		GROUP BY fecha
		ORDER BY fecha DESC
	`, fechas, zona).Scan(&result).Error
	return result, err
}

// Total para la penúltima fecha
func GetYesterdayVisitas(zona string) ([]YesterdayVisitas, error) {
	var fecha string
	err := db.DB.Raw(`
		SELECT DISTINCT fecha 
		FROM visitas 
		WHERE zona = ?
		ORDER BY fecha DESC
		LIMIT 1 OFFSET 1
	`, zona).Scan(&fecha).Error
	if err != nil || fecha == "" {
		return nil, err
	}

	var result []YesterdayVisitas
	err = db.DB.Raw(`
		SELECT 
			fecha,
			zona,
			SUM(visitantes) as total
		FROM visitas
		WHERE fecha = ? AND zona = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16
		GROUP BY fecha, zona
	`, fecha, zona).Scan(&result).Error

	return result, err
}

// Suma por hora en una fecha (ojiva)
func GetOjivaVisitas(fecha, zona string) ([]OjivaVisitas, error) {
	var result []OjivaVisitas

	if fecha == "" {
		err := db.DB.Raw(`SELECT MAX(fecha) FROM visitas WHERE zona = ?`, zona).Scan(&fecha).Error
		if err != nil {
			return result, err
		}
	}

	err := db.DB.Raw(`
		SELECT 
			fecha,
			CONCAT(LPAD(CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED), 2, '0'), ':00') AS hora,
			SUM(visitantes) as total
		FROM visitas
		WHERE fecha = ? AND zona = ? AND CAST(SUBSTRING(hora, 1, 2) AS UNSIGNED) BETWEEN 9 AND 16
		GROUP BY fecha, hora
		ORDER BY hora
	`, fecha, zona).Scan(&result).Error

	return result, err
}
