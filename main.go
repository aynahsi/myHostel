package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofrproject/api/attendances"
	"gofrproject/api/students"
	"log"
	"sync"
)

type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

var (
	clients = make(map[*websocket.Conn]bool) // Map to store clients
	mu      sync.Mutex                       // Mutex to protect clients map
)

func checkWebSocketHeaders(c *gofr.Context) []string {
	missingHeaders := []string{}
	requiredHeaders := []string{"Connection", "Upgrade", "Sec-Websocket-Version", "Sec-WebSocket-Key"}

	for _, key := range requiredHeaders {
		if c.Header(key) == "" {
			missingHeaders = append(missingHeaders, key)
		}
	}

	return missingHeaders
}

func handleDisconnect(conn *websocket.Conn) {
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
	err := conn.Close()
	if err != nil {
		return
	}
}

func broadcastMessage(message Message, sender *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling message: %v", err)
		return
	}

	for client := range clients {
		if client != sender {
			if err := client.WriteMessage(websocket.TextMessage, messageJSON); err != nil {
				handleDisconnect(client)
			}
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := gofr.New()

	// Define your routes and handlers
	app.POST("/student", students.AddStudent)
	app.GET("/student/{studentID}", students.GetStudent)
	app.GET("/students", students.GetAllStudents)
	app.PUT("/student/{studentID}", students.UpdateStudent)
	app.DELETE("/student/{studentID}", students.DeleteStudent)

	app.POST("/attendances", attendances.RecordAttendance)
	app.PUT("/attendances/{recordID}", attendances.UpdateAttendance)
	app.DELETE("/attendances/{recordID}", attendances.DeleteAttendance)
	app.GET("/attendances/{studentID}", attendances.GetAttendance)

	app.GET("/websocket", func(c *gofr.Context) (interface{}, error) {
		conn := c.WebSocketConnection
		if conn != nil {
			mu.Lock()
			clients[conn] = true
			mu.Unlock()

			for {
				_, data, err := conn.ReadMessage()
				if err != nil {
					handleDisconnect(conn)
					break
				}

				var msg Message
				if err := json.Unmarshal(data, &msg); err != nil {
					log.Printf("Error unmarshaling message: %v", err)
					continue
				}

				broadcastMessage(msg, conn)
			}
		} else {
			return nil, errors.MissingParam{Param: checkWebSocketHeaders(c)}
		}

		return nil, nil
	})

	app.Start()
}
