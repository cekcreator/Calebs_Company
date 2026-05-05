package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/calebs-company/realtime/internal/hub"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// TODO: restrict origins in production
		return true
	},
}

type WSHandler struct {
	hub *hub.Hub
}

func NewWSHandler(h *hub.Hub) *WSHandler {
	return &WSHandler{hub: h}
}

func (wsh *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	room := r.URL.Query().Get("room")
	if room == "" {
		room = "global"
	}
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = fmt.Sprintf("anon-%d", time.Now().UnixNano())
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade error: %v", err)
		return
	}

	client := &hub.Client{
		ID:   userID,
		Room: room,
		Send: make(chan []byte, 256),
	}
	wsh.hub.Register(client)

	go writePump(conn, client)
	readPump(conn, client, wsh.hub)
}

func readPump(conn *websocket.Conn, c *hub.Client, h *hub.Hub) {
	defer func() {
		h.Unregister(c)
		conn.Close()
	}()
	conn.SetReadLimit(4096)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		h.Broadcast(c.Room, msg)
	}
}

func writePump(conn *websocket.Conn, c *hub.Client) {
	defer conn.Close()
	for msg := range c.Send {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}
