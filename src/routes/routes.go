package routes

import (
	"api2/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Visitas
	api.GET("/visitas/now", controllers.GetNowVisitas)
	api.GET("/visitas/lastweek", controllers.GetLastWeekVisitas)
	api.GET("/visitas/yesterday", controllers.GetYesterdayVisitas)
	api.GET("/visitas/ojiva", controllers.GetOjivaVisitas)
	api.GET("/visitas/fecha/:fecha", controllers.GetFechaVisitas)

	// Atracciones
	api.GET("/atraccion/now", controllers.GetNowAtraccion)
	api.GET("/atraccion/lastweek", controllers.GetLastWeekAtraccion)
	api.GET("/atraccion/yesterday", controllers.GetYesterdayAtraccion)
	api.GET("/atraccion/ojiva", controllers.GetOjivaAtraccion)
	api.GET("/atraccion/fecha/:fecha", controllers.GetFechaAtraccion)
}
