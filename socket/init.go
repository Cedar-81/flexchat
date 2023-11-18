package socket

import (
	"fmt"
	"net/http"
)

// func (server *WebSocketServer) OnConnect(connection func(*websocket.Conn)) *WebSocketServer {
// 	fmt.Println("i am in here")
// 	go server.processSendQueue()
// 	server.connections = append(server.connections, func(conn *WebSocketConn) {

// 		socket := NewWebSocketConn(conn.Conn)
// 		connection(socket.Conn)
// 		go server.handleEvents(conn)

//		})
//		return server
//	}
func (server *WebSocketServer) handleConnection(writer http.ResponseWriter, request *http.Request) {
	conn, err := server.upgrader.Upgrade(writer, request, nil)
	fmt.Println("chilling here...", conn)
	server.temp_conn = conn

	if err != nil {
		fmt.Println("An error has occurred: ", err)
		return
	}

	fmt.Println("Server up and listening for events...")

	//activate send queue
	go server.processSendQueue()
}

func (server *WebSocketServer) Init(port string) error {

	return http.ListenAndServe(":"+port, nil)
}
