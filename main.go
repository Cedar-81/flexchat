package main

import (
	"fmt"

	"github.com/Cedar-81/flexchat/socket"
	"github.com/gorilla/websocket"
)

func main() {
	server := socket.NewWebSocketServer()

	fmt.Println("Befor init...")

	server.OnConnect(func(conn *websocket.Conn) {

		server.On("join", func(conn *socket.WebSocketConn, data interface{}) {
			server.Join(data, conn)
		})

		server.On("message", func(innerConn *socket.WebSocketConn, data interface{}) {
			fmt.Println("Received 'message' event:", data)
			server.To(*innerConn.Active_room_id, innerConn).Emit("message", data, innerConn)
			// socketInstance.Emit("message", socket.NewEmitOptions(data, data["room_id"]), conn)
		})

	})

	err := server.Init("8080")

	if err != nil {
		fmt.Println("Could not initilaize server: ", err)
	}

}
