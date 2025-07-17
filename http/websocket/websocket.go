package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	port      string
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan []byte)
	mutex     = &sync.Mutex{}
)

func init() {
	flag.StringVar(&port, "a", ":9080", "websocket server address")
}

func main() {
	flag.Parse()
	fmt.Println(port)

	router := gin.Default()
	go broadcastMessage()
	sendMessage(router)

	// Start Server:
	router.Run(port)
}

func conncetionUpgrader(readBuffer int, writeBuffer int, enableCompression bool) websocket.Upgrader {
	upgrader := websocket.Upgrader{
		ReadBufferSize:    readBuffer,
		WriteBufferSize:   writeBuffer,
		EnableCompression: false,
		CheckOrigin: func(r *http.Request) bool {
			// Allow all origins for simplicity, adjust as needed
			return true
		},
	}
	return upgrader
}

func sendMessage(router *gin.Engine) {

	// Make Server GET HTML For WebSocket Client
	router.GET("/", func(c *gin.Context) {
		c.File("./websocket.html")
	})

	// Make Upgrader for WebSocket Connections
	upgrader := conncetionUpgrader(1024, 1024, false)

	// Make get and sent message.
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Error upgrading connection:", err)
			return
		}
		defer conn.Close()

		// Make register connection:
		mutex.Lock()
		clients[conn] = true
		mutex.Unlock()

		for {
			_, message, err := conn.ReadMessage()
			log.Printf("Message is %s", string(message))
			if err != nil {
				log.Println("Cannot read message from websocket ...", err)
				break
			}
			broadcast <- message
		}
	})
}

func broadcastMessage() {
	for {
		message := <-broadcast
		mutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Cannot write message with error :", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
