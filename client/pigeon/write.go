package pigeon

import (
	"fmt"
	"log/slog"

	"github.com/gorilla/websocket"
)

func WriteMessage(inChan chan []string, conn *websocket.Conn) error {

	for {
		msg, ok := <- inChan

		if !ok {
			slog.Error("writeMessage: channel closed")
			return fmt.Errorf("channel closed")
		}

		err := conn.WriteMessage(websocket.TextMessage, []byte(msg[0]))
		if err != nil {
			slog.Error("writeMessage: error writing message", err)
			return err
		}
	}
}