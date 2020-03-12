package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	publicAddress := flag.String("public-address", "foo", "a string")
	port := flag.String("port", "foo", "a string")
	flag.Parse()
	fmt.Println(*publicAddress)
	fmt.Println(*port)

	http.HandleFunc("/connect", connectHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:" + *port, nil))
}


func connectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Printf("%+v\n", r)
	w.Write([]byte("Hi"))
}
