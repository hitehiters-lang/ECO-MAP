package ws

import (
	"eco-map-app/additional"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mu         sync.Mutex
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.register:
			h.mu.Lock()
			h.clients[conn] = true
			h.mu.Unlock()
			additional.PrintSuccess("/ws", fmt.Sprintf("Новый клиент подключен. Всего клиентов: %d", len(h.clients)))

		case conn := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[conn]; ok {
				delete(h.clients, conn)
				conn.Close()
			}
			h.mu.Unlock()
			additional.PrintSuccess("/ws", fmt.Sprintf("Клиент отключен. Всего клиентов: %d", len(h.clients)))

		case message := <-h.broadcast:
			h.mu.Lock()
			for conn := range h.clients {
				go func(c *websocket.Conn) {
					err := c.WriteMessage(websocket.TextMessage, message)
					if err != nil {
						additional.PrintServerError(err.Error())
						c.Close()
						h.unregister <- c
					}
				}(conn)
			}
			h.mu.Unlock()
		}
	}
}

func WsHandler(hub *Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📥 Запрос на WebSocket: %s %s | Upgrade: %s", r.Method, r.URL.Path, r.Header.Get("Upgrade"))
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			additional.PrintServerError(err.Error())
			return
		}

		hub.register <- conn

		go func() {
			defer func() {
				hub.unregister <- conn
				conn.Close()
			}()

			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						additional.PrintServerError(err.Error())
					}
					break
				}
			}
		}()
	}
}
