package controllers

import (
	"api2/src/entities"
	"api2/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateVisitaGeneral godoc
// @Summary Crear una nueva visita
// @Tags visitasGeneral
// @Accept json
// @Produce json
// @Param visita body entities.VisitaGeneral true "Datos de la visita"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/visitasGeneral [post]
func CreateVisitaGeneral(c *gin.Context) {
	var visita entities.VisitaGeneral
	if err := c.ShouldBindJSON(&visita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateVisitaGeneral(visita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Visita creada"})
}

// GetAllVisitasGeneral godoc
// @Summary Obtener todas las visitas registradas
// @Tags visitasGeneral
// @Produce json
// @Success 200 {array} entities.VisitaGeneral
// @Failure 500 {object} map[string]string
// @Router /api/visitasGeneral [get]
func GetAllVisitasGeneral(c *gin.Context) {
	data, err := models.GetAllVisitasGeneral()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetVisitaGeneralByID godoc
// @Summary Obtener una visita por ID
// @Tags visitasGeneral
// @Produce json
// @Param id path int true "ID de la visita"
// @Success 200 {object} entities.VisitaGeneral
// @Failure 404 {object} map[string]string
// @Router /api/visitasGeneral/{id} [get]
func GetVisitaGeneralByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	visita, err := models.GetVisitaGeneralByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Visita no encontrada"})
		return
	}
	c.JSON(http.StatusOK, visita)
}

// UpdateVisitaGeneral godoc
// @Summary Actualizar una visita por ID
// @Tags visitasGeneral
// @Accept json
// @Produce json
// @Param id path int true "ID de la visita"
// @Param visita body entities.VisitaGeneral true "Nuevos datos de la visita"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/visitasGeneral/{id} [put]
func UpdateVisitaGeneral(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updated entities.VisitaGeneral
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateVisitaGeneral(id, updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Visita actualizada"})
}

// DeleteVisitaGeneral godoc
// @Summary Eliminar una visita por ID
// @Tags visitasGeneral
// @Produce json
// @Param id path int true "ID de la visita"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/visitasGeneral/{id} [delete]
func DeleteVisitaGeneral(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := models.DeleteVisitaGeneral(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Visita eliminada"})
}
