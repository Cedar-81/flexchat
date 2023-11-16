package socket

import (
	"fmt"
	"strings"
	"sync"
)

func (server *WebSocketServer) processSendQueue() {
	for msg := range server.sendQueue {
		if msg.room != nil {
			fmt.Println("Processing message: ", msg)
			server.sendMessageToRoom(*msg.room, &msg)
		}
		if msg.room == nil {
			err := msg.conn.Conn.WriteJSON(msg.data)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func (server *WebSocketServer) sendMessageToRoom(id string, msg *Message) {
	var wg sync.WaitGroup
	room, exists := server.rooms[strings.TrimSpace(id)]
	fmt.Println("Sending message to: ", room, exists)
	if exists {
		wg.Add(1)
		go server.broadcastMessageToRoom(room, &wg)
		room.Channel <- msg
	}

}

func (server *WebSocketServer) broadcastMessageToRoom(room *Room, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Broadcasting message to channel: ", room.ID)
	for msg := range room.Channel {
		fmt.Println("Broadcasting... ", msg.data)
		for _, member := range room.Members {
			if strings.TrimSpace(member.id) == strings.TrimSpace(msg.conn.id) {
				continue
			}
			fmt.Println("Sending message to members: ", member.id)
			member.Conn.WriteJSON(msg.data)
		}
	}
}

// func (socket *WebSocket) listenForEvents(conn *WebSocketConn) {
// 	for {
// 		var event map[string]interface{}
// 		err := conn.conn.ReadJSON(&event)

// 		if err!= nil {
//             fmt.Println(err)
//             return
//         }

// 		socket.On()

// 	}
// }

func (server *WebSocketServer) handleEvents(conn *WebSocketConn) {
	fmt.Println("Checking events...")
	for {
		var event map[string]interface{}
		if err := conn.Conn.ReadJSON(&event); err != nil {
			fmt.Println(err)
			return
		}
		for _, handler := range server.eventHandlers[event["type"].(string)] {
			go handleEvent(event, handler, conn)

		}
	}

}

func handleEvent(data map[string]interface{}, handler func(*WebSocketConn, interface{}), connection *WebSocketConn) {
	fmt.Println("handling...")
	fmt.Println("Handling event...", data["type"].(string), data["value"].(string))
	go handler(connection, data["value"])
}
