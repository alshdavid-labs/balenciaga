package websockets

import (
	"fmt"
	"log"
	"os"
	"sync"
	"syscall"

	"github.com/gorilla/websocket"
	"gopkg.in/vrecan/death.v3"
)

// Websocket represents a websocket connection
type Websocket struct {
	Privilaged     bool
	RegisteredName string
	SocketID       string
	Conn           *websocket.Conn
}

func (w *Websocket) SendAndWait(msg []byte) []byte {
	result := []byte{}
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Println(1)
		for {
			_, msg, _ := w.Conn.ReadMessage()
			result = msg
			fmt.Println(2)
			break
		}
		fmt.Println(3)
		wg.Done()
	}()

	w.Send(msg)

	wg.Wait()
	return result
}

// OnMessage is a listener for websocket messages
func (w *Websocket) OnMessage(cb func([]byte)) {
	done := make(chan struct{})
	alive := true

	ded := death.NewDeath(syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer close(done)
		for alive == true {
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

// Send will dispatch a message over
// the websocket connection
func (w *Websocket) Send(msg []byte) {
	err := w.Conn.WriteMessage(
		websocket.TextMessage,
		msg,
	)
	if err != nil {
		return
	}
}

// Close will terminate a websocket connection
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
