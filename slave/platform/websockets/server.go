package websockets

import "net/http"

import "github.com/gorilla/websocket"

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
	return &Websocket{
		Conn: c,
	}
}
