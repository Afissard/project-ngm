package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:4000", Path: "/ws"}
	fmt.Printf("Connecting to %s...\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	// Send "hello"
	err = conn.WriteMessage(websocket.TextMessage, []byte("hello"))
	if err != nil {
		log.Fatal("Write error:", err)
	}
	fmt.Println("Sent: hello")

	// Set read deadline (optional but recommended)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	// Wait for response
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("Read error:", err)
	}
	fmt.Printf("Received: %s\n", message)
}
