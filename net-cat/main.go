package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

var (
	clients []net.Conn
	mutex   sync.Mutex
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening on port 8080")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error adding users")
		}
		mutex.Lock()
		clients = append(clients, conn)
		mutex.Unlock()
		go handleInput(conn)
	}
}

func handleInput(conn net.Conn) {
	message := make([]byte, 4096)
	defer conn.Close()

	for {
		size, err := conn.Read(message)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}

			mutex.Lock()
			for i, client := range clients {
				if client == conn {
					clients = append(clients[:i], clients[i+1:]...)
					fmt.Printf("%v left the chat room", client.RemoteAddr())
				}
			}
			mutex.Unlock()
			break
		}
		fmt.Println(string(message[:size]))
		broadcastMessage(message[:size], conn)
	}
}

func broadcastMessage(msg []byte, sender net.Conn) {
	for _, client := range clients {
		if sender != client {
			mutex.Lock()
			client.Write(msg)
			mutex.Unlock()
		}
	}
}
