package socket

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (server *WebSocketServer) OnConnect(connection func(*websocket.Conn)) *WebSocketServer {
	fmt.Println("i am in here")
	go server.processSendQueue()
	server.connections = append(server.connections, func(conn *WebSocketConn) {

		socket := NewWebSocketConn(conn.Conn)
		connection(socket.Conn)
		go server.handleEvents(conn)

	})
	return server
}

func (server *WebSocketServer) Init(port string) error {
	netHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		conn, err := server.upgrader.Upgrade(writer, request, nil)

		if err != nil {
			fmt.Println("An error has occurred: ", err)
			return
		}

		fmt.Println("Server up and listening for events...")

		// Notify all registered OnConnect handlers about the new connection
		for _, handler := range server.connections {
			fmt.Println("in here")
			handler(&WebSocketConn{
				Conn: conn,
				id:   uuid.New().String(),
			})
		}

	})

	return http.ListenAndServe(":"+port, netHandler)
}
