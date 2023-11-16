package socket

import (
	"fmt"
	"log"
	"strings"
)

func (server *WebSocketServer) On(event string, handler func(*WebSocketConn, interface{})) *WebSocketServer {
	server.eventHandlers[event] = append(server.eventHandlers[event], handler)

	return server
}

func (server *WebSocketServer) Join(data interface{}, conn *WebSocketConn) {

	fmt.Println("Conn: ", conn.id)
	fmt.Println("Available rooms: ", server.rooms)

	fmt.Println("Joining...", data, server.rooms[data.(string)])

	value, ok := data.(string)

	if !ok {
		log.Fatal("Couldn't convert interface to string")
	}

	room, exists := server.rooms[strings.TrimSpace(value)]
	if exists {
		fmt.Println("room exists")
		room.Members[conn.id] = conn
		conn.Active_room_id = &value
		fmt.Println("Found and Joined room successfully. ", room.ID, room.Members)
		return
	}

	fmt.Println("Room does not exist. Creating room...")
	room = NewRoom(strings.TrimSpace(value))
	room.Members[conn.id] = conn
	server.rooms[room.ID] = room
	conn.Active_room_id = &room.ID

	fmt.Println("Room created and joined successfully", room.ID)

}

func (server *WebSocketServer) To(id string, conn *WebSocketConn) *WebSocketServer {
	_, exists := server.rooms[id]
	if !exists {
		fmt.Println("Room doesn't exist")

	}
	conn.Active_room_id = &id

	return server
}

func (server *WebSocketServer) Emit(event string, data interface{}, conn *WebSocketConn) {
	server.sendQueue <- Message{conn: conn, data: map[string]interface{}{"type": event, "data": data}, room: conn.Active_room_id}
}
