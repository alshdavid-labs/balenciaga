package main

import (
	"fmt"
	"parkour/slave/platform/websockets"
)

func main() {
	ws := websockets.Connect("ws://localhost:8080/echo")
	defer ws.Close()

	ws.Send([]byte("Hi"))

	ws.OnMessage(func(msg []byte) {
		fmt.Println(string(msg))
	})
}
