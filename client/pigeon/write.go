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
		
		for i := 0; i < len(msg); i++ {
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg[0]))
			if err != nil {
				slog.Error(fmt.Sprintf("writeMessage: error writing message: %v", err))
				return err
			}
			slog.Info(fmt.Sprintf("writeMessage: sent:\t%s", msg[0]))
		}
		
	}
}