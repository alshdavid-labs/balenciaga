package main

import (
	"flag"
	"log"
	"net/http"

	"parkour/slave/platform/websockets"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	http.HandleFunc("/child", childHandler)

	http.HandleFunc("/connect", connectHandler)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func childHandler(w http.ResponseWriter, r *http.Request) {
	ws := websockets.Upgrade(w, r)
	defer ws.Close()

	ws.OnMessage(func(msg []byte) {
		log.Println(string(msg))

		ws.Send(msg)
	})
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	ws := websockets.Upgrade(w, r)
	defer ws.Close()

	ws.OnMessage(func(msg []byte) {
		log.Println(string(msg))

		ws.Send(msg)
	})
}
