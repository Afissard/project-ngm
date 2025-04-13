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
			slog.Error(fmt.Sprintf("readMessage: error reading message: %v", err))
			return err
		}

		msgReceived := string(msg[:])
		outChan <- []string{msgReceived}
		slog.Info(fmt.Sprintf("readMessage: received:\t%s", msgReceived))
	}
}
