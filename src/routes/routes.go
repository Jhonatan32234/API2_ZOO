package routes

import (
	"api2/src/controllers"
	"api2/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Visitas
	api.GET("/visitas/now", auth.JWTQueryMiddleware("admin", "user"), controllers.GetNowVisitas)
	api.GET("/visitas/lastweek", auth.JWTQueryMiddleware("admin", "user"), controllers.GetLastWeekVisitas)
	api.GET("/visitas/yesterday", auth.JWTQueryMiddleware("admin", "user"), controllers.GetYesterdayVisitas)
	api.GET("/visitas/ojiva", auth.JWTQueryMiddleware("admin", "user"), controllers.GetOjivaVisitas)
	api.GET("/visitas/fecha/:fecha", auth.JWTQueryMiddleware("admin", "user"), controllers.GetFechaVisitas)

	// Atracciones
	api.GET("/atraccion/now", auth.JWTQueryMiddleware("admin", "user"), controllers.GetNowAtraccion)
	api.GET("/atraccion/lastweek", auth.JWTQueryMiddleware("admin", "user"), controllers.GetLastWeekAtraccion)
	api.GET("/atraccion/yesterday", auth.JWTQueryMiddleware("admin", "user"), controllers.GetYesterdayAtraccion)
	api.GET("/atraccion/ojiva", auth.JWTQueryMiddleware("admin", "user"), controllers.GetOjivaAtraccion)
	api.GET("/atraccion/fecha/:fecha", auth.JWTQueryMiddleware("admin", "user"), controllers.GetFechaAtraccion)

	// Visitas General
	api.GET("/visitasGeneral", auth.JWTQueryMiddleware("admin"), controllers.GetAllVisitasGeneral)
	api.GET("/visitasGeneral/:id", auth.JWTQueryMiddleware("admin"), controllers.GetVisitaGeneralByID)
	api.POST("/visitasGeneral", auth.JWTQueryMiddleware("admin"), controllers.CreateVisitaGeneral)
	api.PUT("/visitasGeneral/:id", auth.JWTQueryMiddleware("admin"), controllers.UpdateVisitaGeneral)
	api.DELETE("/visitasGeneral/:id", auth.JWTQueryMiddleware("admin"), controllers.DeleteVisitaGeneral)
}



