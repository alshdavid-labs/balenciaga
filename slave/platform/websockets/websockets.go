package websockets

import (
	"io"
)

type MessageSender interface {
	Send(msg []byte)
}

type MessageListener interface {
	OnMessage(cb func([]byte))
}

type Connector interface {
	MessageSender
	MessageListener
	io.Closer
}
