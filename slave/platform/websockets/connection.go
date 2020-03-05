package websockets

import (
	"log"
	"os"
	"syscall"

	"github.com/gorilla/websocket"
	"gopkg.in/vrecan/death.v3"
)

type Websocket struct {
	Conn *websocket.Conn
}

func (w *Websocket) OnMessage(cb func([]byte)) {
	done := make(chan struct{})

	ded := death.NewDeath(syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer close(done)
		for {
			_, message, err := w.Conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			cb(message)
		}
	}()

	ded.WaitForDeath(w)
	os.Exit(0)
}

func (w *Websocket) Send(msg []byte) {
	err := w.Conn.WriteMessage(
		websocket.TextMessage,
		msg,
	)
	if err != nil {
		return
	}
}

func (w *Websocket) Close() error {
	err := w.Conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)
	if err != nil {
		return err
	}
	w.Conn.Close()
	return nil
}
