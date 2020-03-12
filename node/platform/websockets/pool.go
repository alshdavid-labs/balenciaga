package websockets

import "fmt"

type WebsocketMap map[string]*Websocket

type Pool struct {
	Connections WebsocketMap
}

func CreatePool() *Pool {
	return &Pool{
		Connections: map[string]*Websocket{},
	}
}

func (p *Pool) Add(conn *Websocket) {
	p.Connections[conn.SocketID] = conn
}

func (p *Pool) Remove(conn *Websocket) {
	delete(p.Connections, conn.SocketID)
}

func (p *Pool) GetPrivilaged() WebsocketMap {
	result := WebsocketMap{}
	fmt.Println(p.Connections)
	for _, _conn := range p.Connections {
		conn := _conn
		if conn.Privilaged == true {
			fmt.Println(conn)
			result[conn.SocketID] = conn
		}
	}
	return result
}

func (p *Pool) GetRandomPrivilaged() *Websocket {
	var result *Websocket
	for _, conn := range p.GetPrivilaged() {
		result = conn
	}
	return result
}
