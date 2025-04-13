package pigeon

import (
	"fmt"
	"log/slog"
	"net/url"

	"github.com/gorilla/websocket"
)


func StartConnection(host string, inChan, outChan chan []string) error {
	
	// Set default host to localhost if not provided
	if host == "" {	
		host = "localhost:4000"
	}

	u := url.URL{Scheme: "ws", Host: host, Path: "/ws"}
	slog.Info(fmt.Sprintf("Connecting to %s\n", u.String()))

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		slog.Error("Dial error:", err)
		return nil
	}
	defer conn.Close()

	go WriteMessage(inChan, conn)
	ReadMessage(outChan, conn)

	return nil
}
