package topic

import "parkour/node/platform/websockets"

type Topic struct {
	ID         string
	Privilaged bool
	Conn       *websockets.Websocket
}
