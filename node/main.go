package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"parkour/node/platform/message"
	"parkour/node/platform/topic"
	"parkour/node/platform/websockets"
)

var topics = topic.NewContainer()
var addr = *flag.String("addr", "localhost:8081", "HTTP service address")
var secret = *flag.String("secret", "", "Secret for authorizing privilaged client")
var connectionPool = websockets.CreatePool()

func main() {
	http.HandleFunc("/", connectHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	conn := websockets.Upgrade(w, r)
	connectionPool.Add(conn)

	defer func() {
		conn.Close()
		topics.Remove(conn.SocketID)
	}()

	conn.OnMessage(func(raw []byte) {
		msg := message.FromBytes(raw)

		if msg.Type == message.MessageType.Register {
			conn.RegisteredName = msg.Value.(string)
			return
		}

		if msg.Type == message.MessageType.Subscribe {
			var t string = msg.Value.(string)
			topics.Add(t, conn)

			priv := connectionPool.GetRandomPrivilaged()
			fmt.Println(priv)

			rawResp := priv.SendAndWait(
				message.Create(
					message.MessageType.Challenge,
					conn.RegisteredName,
				).ToBytes())

			resp := message.FromBytes(rawResp)
			if resp.Value.(string) != "SUCCESS" {
				fmt.Println(conn.SocketID, "Failed Subscribed", t)
				return
			}

			fmt.Println(conn.SocketID, "Subscribed", t)
			conn.Send(
				message.
					CreateSubscribed(t).
					ToBytes(),
			)
			return
		}

		if msg.Type == message.MessageType.Authorize {
			if msg.Value.(string) == secret {
				conn.Privilaged = true
				fmt.Println("Authorized S")
				connectionPool.GetPrivilaged()
			}
		}

		if topics.HasKey(msg.Type) == false {
			return
		}

		for _, c := range topics.Get(msg.Type) {
			resp := message.Create(msg.Type, msg.Value)
			c.Send(resp.ToBytes())
		}
	})
}
