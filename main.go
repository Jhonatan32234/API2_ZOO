package main

import (
	"api2/db"
	"api2/src/controllers"
	"api2/src/models/services"
	"api2/src/routes"
	"api2/websocket"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	services.StartRabbitConsumers()
	go websocket.StartBroadcaster()

	r := gin.Default()

	// ✅ Agrega aquí tu ruta WebSocket
	r.GET("/ws", controllers.WebSocketHandler)

	// ✅ Agrega las demás rutas
	routes.SetupRoutes(r)

	log.Println("Servidor iniciado en :8081")
	r.Run(":8081")
}
