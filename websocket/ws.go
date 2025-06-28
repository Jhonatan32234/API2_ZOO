package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"net/http"
	"api2/auth"
	"sync"
	"log"
	"time"
)

var (
	clients   = make(map[*websocket.Conn]string)
	broadcast = make(chan []byte)
	mutex     = &sync.Mutex{}
)

func HandleConnections(c *gin.Context) {
	claims, err := auth.ValidateTokenFromQuery(c, "admin", "dev")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	log.Printf("✅ Usuario autenticado: ID=%d, Rol=%s", claims.UserID, claims.Role)

	upgrader := websocket.Upgrader{
		CheckOrigin:      func(r *http.Request) bool { return true },
		HandshakeTimeout: 30 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("❌ WebSocket upgrade error: %v", err)
		return
	}
	defer func() {
		mutex.Lock()
		delete(clients, ws)
		mutex.Unlock()
		ws.Close()
	}()

	mutex.Lock()
	clients[ws] = c.Query("token")
	mutex.Unlock()

	log.Printf("🟢 Conexión WebSocket establecida.")

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("🔴 Conexión cerrada: %v", err)
			break
		}
		log.Printf("📩 Mensaje recibido: %s", msg)
	}
}

func NotifyClients(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	broadcast <- bytes
}


func StartBroadcaster() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("Error al enviar mensaje: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
