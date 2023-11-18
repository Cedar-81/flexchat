package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func listenForEvents(conn *websocket.Conn) {
	for {
		var message map[string]interface{}
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Fatal("There seems to be an issue", err)
		}

		switch eventType := message["type"]; eventType {
		case "message":
			fmt.Println("Server says: ", message["data"])
		default:
			fmt.Println("Server says nothing")
		}
	}
}

func client() {
	serverAddr := "ws://localhost:8080/ws"
	socketURL, err := url.Parse(serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	conn, _, err := websocket.DefaultDialer.Dial(socketURL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go listenForEvents(conn)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter room id to join>> ")
	room, err := reader.ReadString('\n')
	err = conn.WriteJSON(map[string]interface{}{
		"type":  "join",
		"value": room,
	})

	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("Enter message(enter exit to leave chat)>> ")
		message, err := reader.ReadString('\n')
		if strings.TrimSpace(message) == "exit" {
			break
		}

		fmt.Println("Enter room id(enter exit to leave chat)>> ")
		room, err := reader.ReadString('\n')
		if strings.TrimSpace(room) == "exit" {
			break
		}

		go listenForEvents(conn)

		err = conn.WriteJSON(map[string]interface{}{
			"type":    "message",
			"value":   message,
			"room_id": room,
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Message sent:", message)
	}
}

func main() {
	client()
}
