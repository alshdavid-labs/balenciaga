package websockets

import (
	"io"
)

// MessageSender sends
type MessageSender interface {
	Send(msg []byte)
}

// MessageListener listens
type MessageListener interface {
	OnMessage(cb func([]byte))
}

// Connector is a whole websocket connection
type Connector interface {
	MessageSender
	MessageListener
	io.Closer
}
