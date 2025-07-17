package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "a", ":9080", "websocket server address")
}

func main() {
	flag.Parse()
	fmt.Println(port)

	router := gin.Default()
	makeConnection(router)

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

func makeConnection(router *gin.Engine) {
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
		for {
			messageType, p, err := conn.ReadMessage()
			log.Printf("Message is %s", string(p))
			if err != nil {
				log.Println("Cannot read message from websocket ...", err)
				break
			}
			err = conn.WriteMessage(messageType, p)
			if err != nil {
				log.Println("Cannot write message with error :", err)
			}

		}
	})
}
