package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type Client struct {
	name string
	conn net.Conn
}

var (
	ConvoLog []string
	Clients  []Client
	Mutex    sync.Mutex
)

func main() {
	port := argsHandler(os.Args)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error listening on port %s\n", port)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error adding user")
			continue
		}
		if len(Clients) < 10 {
			go handleConnection(conn)
		} else {
			conn.Write([]byte("This room chat accepts only 10 members."))
			conn.Close()
			log.Printf("Someone failed to access the room")
		}
	}
}

func argsHandler(args []string) string {
	port := "8989"
	if len(args) > 1 {
		if len(args) != 2 {
			fmt.Println("[USAGE]: ./TCPChat $port")
			os.Exit(1)
		}
		port = args[1]
	}
	fmt.Printf("Listening on the port %s\n", port)
	return ":" + port
}

func handleConnection(conn net.Conn) {
	username := loginRequest(conn)
	if username == "" {
		return
	}

	Mutex.Lock()
	newClient := Client{
		name: username,
		conn: conn,
	}
	Clients = append(Clients, newClient)
	Mutex.Unlock()

	userJoined(newClient.name, newClient)
	restorePrevMessage(newClient)

	go handleUserInput(newClient)
}

func loginRequest(conn net.Conn) string {
	conn.Write([]byte(Penguin + "[ENTER YOUR NAME]: "))
	nameReader := bufio.NewReader(conn)
	var username string

	for {
		name, err := nameReader.ReadString('\n')
		if err != nil {
			conn.Write([]byte("Error reading input. Connection closed.\n"))
			conn.Close()
			break
		}

		username = strings.TrimSpace(name)
		if username == "" || !isValidText(username) {
			conn.Write([]byte("Name cannot be empty or contains non-alphanumeric charcters.\n[ENTER YOUR NAME]: "))
			continue
		}
		Mutex.Lock()
		for _, c := range Clients {
			if c.name == username {
				conn.Write([]byte("Name is taken.\n[ENTER YOUR NAME]: "))
				Mutex.Unlock()
				continue
			}
		}
		Mutex.Unlock()
		break
	}
	return strings.TrimSpace(username)
}

func isValidText(text string) bool {
	for _, char := range text {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '\x1b' || char == ' ') {
			return false
		}
	}
	return true
}

func userJoined(username string, client Client) {
	announcement := fmt.Sprintf("%s have joined the room\n", username)
	broadcastAnouncMessage((announcement), client)
}

func handleUserInput(client Client) {
	defer client.conn.Close()

	reader := bufio.NewReader(client.conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			removeUser(client.conn)
			break
		}

		msg = strings.TrimSpace(msg)
		if msg == "" || !isValidText(msg) {
			continue
		}
		fmt.Println("->" + msg)

		broadcastMessage(msg, client)
		refreshInput(client, msg)
	}
}

func refreshInput(client Client, msg string) {
	client.conn.Write([]byte("\033[A\033[K"))
	client.conn.Write([]byte(formatMessage(client.name, msg)))
}

func removeUser(conn net.Conn) {
	Mutex.Lock()
	for i, client := range Clients {
		if client.conn == conn {
			Clients = append(Clients[:i], Clients[i+1:]...)
			leftMessage := fmt.Sprintf("%v left the chat room\n", client.name)
			broadcastAnouncMessage((leftMessage), client)
		}
	}
	defer Mutex.Unlock()
}

func broadcastAnouncMessage(msg string, sender Client) {
	formattedMessage := fmt.Sprintf("-> %s ", string(msg))
	for _, client := range Clients {
		if sender.conn != client.conn {
			_, err := client.conn.Write([]byte(formattedMessage))
			if err != nil {
				fmt.Println("Error writing to", client.name, ":", err)
			}
		}
	}
}

func broadcastMessage(msg string, sender Client) {
	formattedMessage := formatMessage(sender.name, msg)
	saveMessage(string(formattedMessage))
	for _, client := range Clients {
		if sender.conn != client.conn {
			_, err := client.conn.Write(formattedMessage)
			if err != nil {
				fmt.Println("Error writing to", client.name, ":", err)
			}
		}
	}
}

func formatMessage(username, message string) []byte {
	time := time.Now().Format("2006-01-02 15:04:05")
	return []byte(fmt.Sprintf("[%s][%s]: %s\n", time, username, message))
}

func saveMessage(msg string) {
	ConvoLog = append(ConvoLog, msg)
}

func restorePrevMessage(client Client) {
	_, err := client.conn.Write([]byte(strings.Join(ConvoLog, "")))
	if err != nil {
		fmt.Println("Error writing to", client.name, ":", err)
	}
}
