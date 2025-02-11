package chat

import (
	"fmt"
	"sync"
)

type Client struct {
	ID      string
	Message chan string
	History []string
}

type ChatRoom struct {
	clients   map[string]*Client
	mu        sync.Mutex
	broadcast chan string
}

func NewChatRoom() *ChatRoom {
	room := &ChatRoom{
		clients:   make(map[string]*Client),
		broadcast: make(chan string, 100),
	}
	go room.start()
	return room
}

func (cr *ChatRoom) start() {
	for msg := range cr.broadcast {
		cr.mu.Lock()
		clientsCopy := make([]*Client, 0, len(cr.clients))
		for _, client := range cr.clients {
			clientsCopy = append(clientsCopy, client)
		}
		cr.mu.Unlock()

		for _, client := range clientsCopy {
			select {
			case client.Message <- msg:
				fmt.Println(" Delivered message to:", client.ID)
			default:
				fmt.Println(" Skipping message for", client.ID, "as buffer is full")
			}
		}
	}
}

func (cr *ChatRoom) Join(clientID string) *Client {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	client := &Client{
		ID:      clientID,
		Message: make(chan string, 100),
		History: []string{},
	}

	cr.clients[clientID] = client

	// Send past messages to the new client
	for _, msg := range client.History {
		client.Message <- msg
	}

	return client
}

func (cr *ChatRoom) Send(clientID, message string) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	fullMessage := clientID + ": " + message
	fmt.Println(" Broadcasting:", fullMessage)

	for _, client := range cr.clients {
		client.History = append(client.History, fullMessage)

		select {
		case client.Message <- fullMessage:
			fmt.Println(" Message sent to", client.ID)
		default:
			fmt.Println(" Message buffer full for", client.ID, ", skipping...")
		}
	}
}

func (cr *ChatRoom) Leave(clientID string) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	if client, exists := cr.clients[clientID]; exists {
		close(client.Message)
		delete(cr.clients, clientID)
		fmt.Println(" Client left:", clientID)
	} else {
		fmt.Println(" Client", clientID, "not found.")
	}
}

func (cr *ChatRoom) GetMessages(clientID string) []string {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	client, exists := cr.clients[clientID]
	if !exists {
		fmt.Println(" Client not found:", clientID)
		return nil
	}

	messages := client.History
	client.History = []string{}
	return messages
}
