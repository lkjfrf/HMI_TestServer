package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

// OutputTorque matches the C++ struct
type OutputTorque struct {
	State     uint8   // BYTE bState
	Mode      uint8   // BYTE bMode
	Pulse     float32 // float fPulse
	Velocity  float32 // float fVelocity
	Angle     float32 // float fAngle
	OutTorque float32 // float fOutTorque
	InTorque  float32 // float fInTorque
}

func main() {
	// Connect to INNO program's port
	conn, err := net.Dial("tcp", "localhost:1998")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Create buffer to receive data
	buffer := make([]byte, 1024)

	for {
		// Read incoming data
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		if n > 0 {
			// Convert binary data to struct
			data := OutputTorque{}
			reader := bytes.NewReader(buffer[:n])
			if err := binary.Read(reader, binary.LittleEndian, &data); err != nil {
				fmt.Println("Error decoding data:", err)
				continue
			}

			// Print received data
			fmt.Printf("Received Data:\n")
			fmt.Printf("State: %d\n", data.State)
			fmt.Printf("Mode: %d\n", data.Mode)
			fmt.Printf("Pulse: %.2f\n", data.Pulse)
			fmt.Printf("Velocity: %.2f\n", data.Velocity)
			fmt.Printf("Angle: %.2f\n", data.Angle)
			fmt.Printf("OutTorque: %.2f\n", data.OutTorque)
			fmt.Printf("InTorque: %.2f\n", data.InTorque)
		}
	}
}
