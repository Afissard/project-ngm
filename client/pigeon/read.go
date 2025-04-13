package pigeon

import (
	"fmt"
	"log/slog"

	"github.com/gorilla/websocket"
)

func ReadMessage(outChan chan []string, conn *websocket.Conn) error {

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			slog.Error("readMessage: error reading message", err)
			return err
		}

		msgReceived := string(msg[:])
		outChan <- []string{msgReceived}
		slog.Info(fmt.Sprintf("readMessage: received:\t%s", msgReceived))
	}
}
