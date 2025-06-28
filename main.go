package main

import (
	"api2/db"
	"api2/src/controllers"
	"api2/src/models/services"
	"api2/src/routes"
	"api2/websocket"
	"log"
	_ "api2/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	db.Connect()

	services.StartRabbitConsumers()
	go websocket.StartBroadcaster()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ws", controllers.WebSocketHandler)

	routes.SetupRoutes(r)

	log.Println("Servidor iniciado en :8081")
	r.Run(":8081")
}
