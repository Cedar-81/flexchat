package socket

import (
	"sync"

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
	server         *WebSocketServer
}

type WebSocketServer struct {
	temp_conn     *websocket.Conn
	connections   []*WebSocketConn
	rooms         map[string]*Room
	upgrader      websocket.Upgrader
	sendQueue     chan Message
	eventHandlers map[string][]func(interface{})
	WaitGroup     sync.WaitGroup
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		upgrader:      websocket.Upgrader{},
		rooms:         make(map[string]*Room),
		sendQueue:     make(chan Message),
		eventHandlers: make(map[string][]func(interface{})),
	}
}

func NewWebSocketConn(conn *websocket.Conn, server_conn *WebSocketServer) *WebSocketConn {
	return &WebSocketConn{
		id:     uuid.New().String(),
		Conn:   conn,
		server: server_conn,
	}
}
