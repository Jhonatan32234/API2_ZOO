package controllers

import (
	"api2/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNowVisitas(c *gin.Context) {
	data, err := models.GetNowVisitas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetLastWeekVisitas(c *gin.Context) {
	data, err := models.GetLastWeekVisitas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetYesterdayVisitas(c *gin.Context) {
	data, err := models.GetYesterdayVisitas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetOjivaVisitas(c *gin.Context) {
	fecha := c.Query("fecha")
	data, err := models.GetOjivaVisitas(fecha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetFechaVisitas(c *gin.Context) {
	fecha := c.Param("fecha")
	data, err := models.GetFechaVisitas(fecha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
