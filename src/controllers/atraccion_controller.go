package controllers

import (
	"api2/src/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNowAtraccion(c *gin.Context) {
	data, err := models.GetNowAtraccion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetLastWeekAtraccion(c *gin.Context) {
	data, err := models.GetLastWeekAtraccion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetYesterdayAtraccion(c *gin.Context) {
	data, err := models.GetYesterdayAtraccion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetOjivaAtraccion(c *gin.Context) {
	fecha := c.Query("fecha")
	if fecha == "" {
		log.Print("Fecha no proporcionada, se usara la actual")
	}
	data, err := models.GetOjivaAtraccion(fecha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetFechaAtraccion(c *gin.Context) {
	fecha := c.Param("fecha")
	data, err := models.GetFechaAtraccion(fecha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
