package main

import (
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
	"gofrproject/api/attendances"
	"gofrproject/api/students"
	"gofrproject/models"
	"log"
	"sync"
)

var (
	clients = make(map[*websocket.Conn]bool) // Connected clients
	mu      sync.Mutex                       // Mutex to protect clients
)

func handleWebSocket(ctx *gofr.Context) (interface{}, error) {
	conn := ctx.WebSocketConnection
	if conn == nil {
		ctx.Logger.Error("Not a WebSocket request")
		return nil, nil
	}

	// Register new client
	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	// Handle incoming messages
	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			handleDisconnect(conn)
			break
		}

		broadcastMessage(msg, conn)
	}

	return nil, nil
}

func handleDisconnect(conn *websocket.Conn) {
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
	conn.Close()
}

func broadcastMessage(msg models.Message, sender *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		if client != sender {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
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
	app.GET("/websocket", handleWebSocket)

	app.Start()
}
