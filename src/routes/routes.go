package routes

import (
	"api2/src/controllers"
	"api2/utils"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Visitas
	api.GET("/visitas/now", utils.JWTQueryMiddleware("admin", "user"), controllers.GetNowVisitas)
	api.GET("/visitas/lastweek", utils.JWTQueryMiddleware("admin", "user"), controllers.GetLastWeekVisitas)
	api.GET("/visitas/yesterday", utils.JWTQueryMiddleware("admin", "user"), controllers.GetYesterdayVisitas)
	api.GET("/visitas/ojiva", utils.JWTQueryMiddleware("admin", "user"), controllers.GetOjivaVisitas)
	api.GET("/visitas/fecha/:fecha", utils.JWTQueryMiddleware("admin", "user"), controllers.GetFechaVisitas)

	// Atracciones
	api.GET("/atraccion/now", utils.JWTQueryMiddleware("admin", "user"), controllers.GetNowAtraccion)
	api.GET("/atraccion/lastweek", utils.JWTQueryMiddleware("admin", "user"), controllers.GetLastWeekAtraccion)
	api.GET("/atraccion/yesterday", utils.JWTQueryMiddleware("admin", "user"), controllers.GetYesterdayAtraccion)
	api.GET("/atraccion/ojiva", utils.JWTQueryMiddleware("admin", "user"), controllers.GetOjivaAtraccion)
	api.GET("/atraccion/fecha/:fecha", utils.JWTQueryMiddleware("admin", "user"), controllers.GetFechaAtraccion)

	// Visitas General
	api.GET("/visitasGeneral", utils.JWTQueryMiddleware("admin"), controllers.GetAllVisitasGeneral)
	api.GET("/visitasGeneral/:id", utils.JWTQueryMiddleware("admin"), controllers.GetVisitaGeneralByID)
	api.POST("/visitasGeneral", utils.JWTQueryMiddleware("admin"), controllers.CreateVisitaGeneral)
	api.PUT("/visitasGeneral/:id", utils.JWTQueryMiddleware("admin"), controllers.UpdateVisitaGeneral)
	api.DELETE("/visitasGeneral/:id", utils.JWTQueryMiddleware("admin"), controllers.DeleteVisitaGeneral)
}



