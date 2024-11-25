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
	// 서버에 연결
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// JSON 메시지 전송
	msg := Message{Type: "greeting", Content: "Hello, server!"}
	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(&msg); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// JSON 응답 수신
	var response Message
	decoder := json.NewDecoder(conn)
	if err := decoder.Decode(&response); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("Received response: %v\n", response)
}
