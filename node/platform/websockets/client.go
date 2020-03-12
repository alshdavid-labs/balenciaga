package websockets

import (
	"log"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// Connect will create a new WebSocket Client
func Connect(address string) *Websocket {
	c, _, err := websocket.DefaultDialer.Dial(address, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	ID := uuid.NewV4()
	return &Websocket{
		Privilaged: false,
		SocketID:   ID.String(),
		Conn:       c,
	}
}
