package api

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[*Client]bool
	groups  map[string][]string // groupName -> usernames
	sync.Mutex
}

type Client struct {
	conn     *websocket.Conn
	username string
	//Gropus []int // this should from backend auth
}

type Message struct {
	Type    string   `json:"type"` // "message", "user_list", "connect", "disconnect", "create_group", "group_list"
	From    string   `json:"from"`
	To      string   `json:"to,omitempty"`      // For private or group messages
	Content string   `json:"content,omitempty"` // The actual message
	Users   []string `json:"users,omitempty"`   // For user list or group creation
	Time    string   `json:"time,omitempty"`    // Timestamp
}

var hub = &Hub{clients: make(map[*Client]bool), groups: make(map[string][]string)}

func (h *Hub) addClient(client *Client) {
	h.Lock()
	defer h.Unlock()

	h.clients[client] = true
}

func (h *Hub) removeClient(client *Client) {
	h.Lock()
	defer h.Unlock()

	delete(h.clients, client)
	client.conn.Close()
}

func (h *Hub) processs(msg Message) {
	h.Lock()
	defer h.Unlock()

	switch msg.Type {
	case "message":
		msg.Time = time.Now().Format("15:04:05")
		if msg.To != "" {
			if users, ok := h.groups[msg.To]; ok {
				for client := range h.clients {
					if contains(users, client.username) {
						h.sendMessage(client, msg)
					}
				}
			} else {
				for client := range h.clients {
					if client.username == msg.To || client.username == msg.From {
						h.sendMessage(client, msg)
					}
				}
			}
		} else {
			for client := range h.clients {
				h.sendMessage(client, msg)
			}
		}
	case "create_group":
		if msg.To != "" && len(msg.Users) > 0 {
			h.groups[msg.To] = append(msg.Users, msg.From)
			h.sendGroupListToMembers(msg.To)
		}
	}
}

func (h *Hub) sendUserList() {
	h.Lock()
	defer h.Unlock()

	var userList []string
	for client := range h.clients {
		userList = append(userList, client.username)
	}

	msg := Message{
		Type:  "user_list",
		Users: userList,
	}

	for client := range h.clients {
		h.sendMessage(client, msg)
	}
}

func (h *Hub) sendGroupList(username string) {
	var userGroups []string
	for groupName, members := range h.groups {
		if contains(members, username) {
			userGroups = append(userGroups, groupName)
		}
	}
	msg := Message{
		Type:  "group_list",
		Users: userGroups,
	}
	for client := range h.clients {
		if client.username == username {
			h.sendMessage(client, msg)
		}
	}
}

func (h *Hub) sendGroupListToMembers(groupName string) {
	members := h.groups[groupName]
	msg := Message{
		Type:  "group_list",
		Users: []string{groupName},
	}
	for client := range h.clients {
		if contains(members, client.username) {
			h.sendMessage(client, msg)
		}
	}
}

func (h *Hub) sendMessage(client *Client, msg Message) {
	err := client.conn.WriteJSON(msg)
	if err != nil {
		log.Printf("error: %v", err)
		client.conn.Close()
		delete(h.clients, client)
	}
}

func contains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
