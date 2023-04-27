package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

type UserConnections map[int]map[string]*websocket.Conn

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	userConnections = make(UserConnections)
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to upgrade connection to WebSocket: %v", err)
		return
	}

	defer conn.Close()

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		log.Printf("invalid user ID: %v", err)
		return
	}

	deviceID := r.URL.Query().Get("device_id")

	if _, ok := userConnections[userID]; !ok {
		userConnections[userID] = make(map[string]*websocket.Conn)
	}

	userConnections[userID][deviceID] = conn

	log.Printf("accepted WebSocket connection from user %d, device %s", userID, deviceID)

	for {
		// read from WebSocket if needed
	}
}

func main() {
	http.HandleFunc("/", handleWebSocket)

	log.Println("Starting WebSocket server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start WebSocket server: %v", err)
	}
}
