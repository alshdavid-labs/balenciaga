package websockets

import (
	"log"

	"github.com/gorilla/websocket"
)

func Connect(address string) *Websocket {
	c, _, err := websocket.DefaultDialer.Dial(address, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return &Websocket{
		Conn: c,
	}
}
