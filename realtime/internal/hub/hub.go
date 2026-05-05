package hub

import (
	"encoding/json"
	"log"
	"sync"
)

type Message struct {
	Type      string          `json:"type"`
	Payload   json.RawMessage `json:"payload"`
	Timestamp string          `json:"timestamp"`
	Room      string          `json:"room,omitempty"`
}

type Client struct {
	ID   string
	Room string
	Send chan []byte
}

type Hub struct {
	mu      sync.RWMutex
	rooms   map[string]map[*Client]struct{}
	register   chan *Client
	unregister chan *Client
	broadcast  chan roomMessage
}

type roomMessage struct {
	room    string
	payload []byte
}

func New() *Hub {
	return &Hub{
		rooms:      make(map[string]map[*Client]struct{}),
		register:   make(chan *Client, 64),
		unregister: make(chan *Client, 64),
		broadcast:  make(chan roomMessage, 256),
	}
}

func (h *Hub) Register(c *Client)   { h.register <- c }
func (h *Hub) Unregister(c *Client) { h.unregister <- c }

func (h *Hub) Broadcast(room string, payload []byte) {
	h.broadcast <- roomMessage{room: room, payload: payload}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.mu.Lock()
			if h.rooms[c.Room] == nil {
				h.rooms[c.Room] = make(map[*Client]struct{})
			}
			h.rooms[c.Room][c] = struct{}{}
			h.mu.Unlock()
			log.Printf("client %s joined room %s", c.ID, c.Room)

		case c := <-h.unregister:
			h.mu.Lock()
			delete(h.rooms[c.Room], c)
			close(c.Send)
			h.mu.Unlock()
			log.Printf("client %s left room %s", c.ID, c.Room)

		case msg := <-h.broadcast:
			h.mu.RLock()
			for c := range h.rooms[msg.room] {
				select {
				case c.Send <- msg.payload:
				default:
					close(c.Send)
					delete(h.rooms[msg.room], c)
				}
			}
			h.mu.RUnlock()
		}
	}
}
