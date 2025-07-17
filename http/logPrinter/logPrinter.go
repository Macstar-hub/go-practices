package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

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
	flag.StringVar(&port, "a", ":9080", "Server Run Port")
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

func sendLines() {
	// Open file from os:
	file, err := os.Open("/Users/Shared/codes.dir/go.dir/git.dir/ti-idf/logs/redis/debug.log")
	if err != nil {
		log.Println("Cannot open file with error: ", err)
	}
	defer file.Close()
	for {
		// Make sleep 1
		time.Sleep(1 * time.Second)

		// Sync Input from file:
		scanner := bufio.NewScanner(file)

		// Read line by line:
		scanner.Split(bufio.ScanLines)

		// Send Line throw channel:
		for scanner.Scan() {
			broadcast <- scanner.Bytes()
		}
	}
}

func sendMessage(router *gin.Engine) {
	// Load client page:
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
		log.Println("Read Line: ", message)
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

func main() {
	// Parse All Flags
	flag.Parse()

	// Set Server Properties
	router := gin.Default()
	go broadcastMessage()
	go sendLines()
	sendMessage(router)

	// Start Server:
	router.Run(port)
}
