package controllers

import (
	"api2/src/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNowAtraccion godoc
// @Summary Obtener atracciones de la fecha más reciente
// @Tags atraccion
// @Produce json
// @Success 200 {array} entities.Atraccion
// @Failure 500 {object} map[string]string
// @Router /api/atraccion/now [get]
func GetNowAtraccion(c *gin.Context) {
	zona := c.GetString("zona")
	data, err := models.GetNowAtraccion(zona)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetLastWeekAtraccion godoc
// @Summary Obtener atracciones de las 6 fechas más recientes
// @Tags atraccion
// @Produce json
// @Success 200 {array} entities.Atraccion
// @Failure 500 {object} map[string]string
// @Router /api/atraccion/lastweek [get]
func GetLastWeekAtraccion(c *gin.Context) {
	zona := c.GetString("zona")
	data, err := models.GetLastWeekAtraccion(zona)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetYesterdayAtraccion godoc
// @Summary Obtener atracciones de la penúltima fecha registrada
// @Tags atraccion
// @Produce json
// @Success 200 {array} entities.Atraccion
// @Failure 500 {object} map[string]string
// @Router /api/atraccion/yesterday [get]
func GetYesterdayAtraccion(c *gin.Context) {
	zona := c.GetString("zona")
	data, err := models.GetYesterdayAtraccion(zona)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetOjivaAtraccion godoc
// @Summary Obtener ojiva de atracción (tiempo total por hora)
// @Tags atraccion
// @Produce json
// @Param fecha query string false "Fecha en formato YYYY-MM-DD"
// @Success 200 {array} models.OjivaResultAtraccion
// @Failure 500 {object} map[string]string
// @Router /api/atraccion/ojiva [get]
func GetOjivaAtraccion(c *gin.Context) {
	zona := c.GetString("zona")
	fecha := c.Query("fecha")
	if fecha == "" {
		log.Print("Fecha no proporcionada, se usara la actual")
	}
	data, err := models.GetOjivaAtraccion(fecha, zona)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

