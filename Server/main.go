package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func main() {
	// 서버 시작
	listener, err := net.Listen("tcp", ":1998")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 1998...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// JSON 메시지 수신
	var msg Message
	decoder := json.NewDecoder(conn)
	if err := decoder.Decode(&msg); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("Received message: %v\n", msg)

	// JSON 응답 전송
	response := Message{Type: "response", Content: "Message received!"}
	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(&response); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}
