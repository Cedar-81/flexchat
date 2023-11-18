package socket

import (
	"fmt"
	"log"
	"strings"
)

func (socket *WebSocketConn) On(event string, handler func(interface{})) *WebSocketConn {
	socket.server.eventHandlers[event] = append(socket.server.eventHandlers[event], handler)

	return socket
}

func (socket *WebSocketConn) Join(data interface{}) {

	fmt.Println("Conn: ", socket.id)
	fmt.Println("Available rooms: ", socket.server.rooms)

	value, ok := data.(map[string]interface{})

	if !ok {
		log.Fatal("Couldn't convert interface to string ", data)
	}

	value_str := strings.TrimSpace(value["value"].(string))

	room, exists := socket.server.rooms[value_str]
	if exists {
		fmt.Println("room exists")
		room.Members[socket.id] = socket
		socket.Active_room_id = &value_str
		fmt.Println("Found and Joined room successfully. ", room.ID, room.Members)
		return
	}

	fmt.Println("Room does not exist. Creating room...")
	room = NewRoom(strings.TrimSpace(value_str))
	room.Members[socket.id] = socket
	socket.server.rooms[room.ID] = room
	socket.Active_room_id = &room.ID

	fmt.Println("Room created and joined successfully", room.ID)

}

func (socket *WebSocketConn) To(value interface{}) *WebSocketConn {
	data, ok := value.(map[string]interface{})
	if !ok {
		log.Fatal("Couldn't convert To id data...", value)
	}
	id := strings.TrimSpace(data["room_id"].(string))
	fmt.Println("To id: ", id)

	room, exists := socket.server.rooms[id]
	if !exists {
		fmt.Println("Room doesn't exist...")
	}

	fmt.Println("Room exists... \nChecking membership...")
	_, exists = room.Members[socket.id]
	if !exists {
		log.Fatal("Sorry, you are not a member of this room.")
	}

	socket.Active_room_id = &id

	return socket
}

func (socket *WebSocketConn) Emit(event string, data interface{}) {
	conv_data, ok := data.(map[string]interface{})

	if !ok {
		log.Fatal("Couldn't convert data in Emitter: ", data)
	}

	value := strings.TrimSpace(conv_data["value"].(string))

	socket.server.sendQueue <- Message{conn: socket, data: map[string]interface{}{"type": event, "data": value}, room: socket.Active_room_id}
}
