package main

import (
	"api2/db"
	_ "api2/docs"
	"api2/src/controllers"
	"api2/src/models/services"
	"api2/src/routes"
	"api2/utils"
	"api2/websocket"
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load(".env")
    if err != nil {
        log.Println("Advertencia: no se pudo cargar el archivo .env:", err)
    }
	db.Connect()

	services.StartRabbitConsumers()
	go websocket.StartBroadcaster()
	go utils.StartBroadcaster()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ws", controllers.WebSocketHandler)

	r.Use(utils.CORSMiddleware())

	routes.SetupRoutes(r)

	log.Println("Servidor iniciado en :8080")
	r.Run(":8080")
}
