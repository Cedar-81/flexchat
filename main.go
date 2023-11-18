package main

import (
	"fmt"

	"github.com/Cedar-81/flexchat/socket"
)

func main() {
	server := socket.NewWebSocketServer()

	fmt.Println("Befor init...")

	server.On("connection", func(socket *socket.WebSocketConn) {

		socket.On("join", func(data interface{}) {
			fmt.Println("joining room....")
			socket.Join(data)
		})

		socket.On("message", func(data interface{}) {
			fmt.Println("Received 'message' event:", data)
			socket.To(data).Emit("message", data)
		})

	})

	// server.WaitGroup.Wait()
	err := server.Init("8080")

	if err != nil {
		fmt.Println("Could not initilaize server: ", err)
	}

}
