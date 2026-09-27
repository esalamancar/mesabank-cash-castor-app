// Package ws implementa el hub de WebSocket: una conexión por cliente,
// agrupadas por código de partida (T-06 del backlog de Sprint 0). Todavía
// sin lógica de negocio — el broadcast de eventos reales (saldo
// actualizado, transacción, arqueo) se agrega al implementar EP-05.
package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Hub mantiene, por código de partida, el conjunto de conexiones activas.
type Hub struct {
	mu    sync.RWMutex
	rooms map[string]map[*websocket.Conn]bool
}

func NewHub() *Hub {
	return &Hub{rooms: make(map[string]map[*websocket.Conn]bool)}
}

func (h *Hub) register(code string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.rooms[code] == nil {
		h.rooms[code] = make(map[*websocket.Conn]bool)
	}
	h.rooms[code][conn] = true
}

func (h *Hub) unregister(code string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.rooms[code], conn)
	if len(h.rooms[code]) == 0 {
		delete(h.rooms, code)
	}
}

// Broadcast envía message a todos los clientes conectados a una partida.
func (h *Hub) Broadcast(code string, message []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for conn := range h.rooms[code] {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("ws: error enviando a partida %s: %v", code, err)
		}
	}
}

var upgrader = websocket.Upgrader{
	// TODO: restringir CheckOrigin antes de producción (RNF-06).
	CheckOrigin: func(r *http.Request) bool { return true },
}

// ServeHTTP hace el upgrade a WebSocket y mantiene la conexión registrada
// en el hub mientras el cliente siga conectado.
func ServeHTTP(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("ws: fallo en upgrade para partida %s: %v", code, err)
			return
		}
		defer conn.Close()

		hub.register(code, conn)
		defer hub.unregister(code, conn)

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				break
			}
		}
	}
}
