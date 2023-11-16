package interfaces

import "github.com/gorilla/websocket"

type WebSocketServer struct {
	Connections []func(*websocket.Conn)
	Upgrader    websocket.Upgrader
}

type Room struct {
	ID      string
	Channel chan string
	Members []*websocket.Conn
}

type WebSocket struct {
	EventHandlers map[string][]func(*websocket.Conn, interface{})
	Room          Room
}
