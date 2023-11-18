package socket

import (
	"fmt"
	"net/http"
	"strings"
)

// servers On function equvalent to io.on
func (server *WebSocketServer) On(event string, handler func(*WebSocketConn)) *WebSocketServer {
	fmt.Println("starting socket...", server.temp_conn)

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := server.upgrader.Upgrade(writer, request, nil)
		fmt.Println("chilling here...", conn)

		if strings.TrimSpace(event) == "connection" {
			fmt.Println("now here...")
			//create new socket connection
			socket := NewWebSocketConn(conn, server)

			//add socket to server connections
			server.connections = append(server.connections, socket)

			//activate event handlers
			// server.WaitGroup.Add(1)
			go server.handleEvents(socket)

			//activate send queue
			go server.processSendQueue()

			//handle connection
			handler(socket)
		}

		if err != nil {
			fmt.Println("An error has occurred: ", err)
			return
		}

		fmt.Println("Server up and listening for events...")

	})

	return server
}
