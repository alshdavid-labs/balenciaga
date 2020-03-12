package websockets

import (
	"net/http"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// Upgrade will upgrade an HTTP request
// to a websocket connection
func Upgrade(
	w http.ResponseWriter,
	r *http.Request,
) *Websocket {
	var u = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	c, err := u.Upgrade(w, r, nil)
	if err != nil {
		return nil
	}
	ID := uuid.NewV4()
	return &Websocket{
		Privilaged: false,
		SocketID:   ID.String(),
		Conn:       c,
	}
}
