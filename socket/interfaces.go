package socket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Room struct {
	ID      string
	Name    string
	Channel chan *Message
	Members map[string]*WebSocketConn
}

type ActiveRoom struct {
	Room *Room
}

func NewRoom(name string) *Room {
	return &Room{
		ID:      uuid.New().String(),
		Name:    name,
		Channel: make(chan *Message),
		Members: make(map[string]*WebSocketConn),
	}
}

// type WebSocketConn struct {
// 	Conn           *websocket.Conn
// 	id             string
// 	Active_room_id *string
// }

type Message struct {
	conn *WebSocketConn
	data interface{}
	room *string
}

type WebSocketConn struct {
	Conn           *websocket.Conn
	id             string
	Active_room_id *string
}

type WebSocketServer struct {
	connections   []func(*WebSocketConn)
	rooms         map[string]*Room
	upgrader      websocket.Upgrader
	sendQueue     chan Message
	eventHandlers map[string][]func(*WebSocketConn, interface{})
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		upgrader:      websocket.Upgrader{},
		rooms:         make(map[string]*Room),
		sendQueue:     make(chan Message),
		eventHandlers: make(map[string][]func(*WebSocketConn, interface{})),
	}
}

func NewWebSocketConn(conn *websocket.Conn) *WebSocketConn {
	return &WebSocketConn{
		id:   uuid.New().String(),
		Conn: conn,
	}
}
