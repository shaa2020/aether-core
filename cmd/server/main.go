package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	fmt.Println("New Peer connected to Aether Core.")

	for {
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Disconnected: %v", err)
			break
		}
		log.Printf("Message Received: %s", msg)

		// Echo back to test the connection
		if err := ws.WriteMessage(messageType, msg); err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	fmt.Println("Aether Core node online on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
