package models

import (
	"api2/db"
	"api2/src/entities"
	"errors"
)

// Crear una nueva visita si no existe ya una con la misma fecha
func CreateVisitaGeneral(visita entities.VisitaGeneral) error {
	var existing entities.VisitaGeneral
	result := db.DB.Where("fecha = ?", visita.Fecha).First(&existing)
	if result.Error == nil {
		return errors.New("ya existe una visita con esa fecha")
	}

	return db.DB.Create(&visita).Error
}

func GetAllVisitasGeneral() ([]entities.VisitaGeneral, error) {
	var visitas []entities.VisitaGeneral
	err := db.DB.Find(&visitas).Error
	return visitas, err
}

func GetVisitaGeneralByID(id int) (entities.VisitaGeneral, error) {
	var visita entities.VisitaGeneral
	err := db.DB.First(&visita, id).Error
	return visita, err
}

func UpdateVisitaGeneral(id int, updated entities.VisitaGeneral) error {
	var visita entities.VisitaGeneral
	err := db.DB.First(&visita, id).Error
	if err != nil {
		return err
	}

	var existing entities.VisitaGeneral
	result := db.DB.Where("fecha = ? AND id != ?", updated.Fecha, id).First(&existing)
	if result.Error == nil {
		return errors.New("ya existe otra visita con esa fecha")
	}

	visita.Fecha = updated.Fecha
	visita.Visitas = updated.Visitas
	return db.DB.Save(&visita).Error
}

func DeleteVisitaGeneral(id int) error {
	return db.DB.Delete(&entities.VisitaGeneral{}, id).Error
}
