package helpers

import "github.com/gorilla/websocket"

type Room struct {
	ID      string
	Channel chan string
	Members []*websocket.Conn
}

var rooms = make(map[string]*Room)

func CreateRoom(id string, member *websocket.Conn) Room {

	room := Room{
		ID:      id,
		Channel: make(chan string),
	}

	room.Members = append(room.Members, member)

	rooms[id] = &room

	return room
}

func GetRoomByID(id string) *Room {
	return rooms[id]
}

func CloseRoom(id string) {
	if room, exists := rooms[id]; exists {
		close(room.Channel)
		delete(rooms, id)
	}

}
