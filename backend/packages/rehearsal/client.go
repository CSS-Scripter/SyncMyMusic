package rehearsal

import (
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
)

// Client struct
type Client struct {
	ID   int
	Conn *websocket.Conn
	Pool *Pool
}

// Message struct
type Message struct {
	Type        int    `json:"type"`
	Body        string `json:"body"`
	RehearsalID int    `json:"rehearsalId"`
	OwnerID     int    `json:"ownerId"`
}

// Read connect client to pool and listen until disconnected
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	var (
		mt  int
		msg []byte
		err error
	)

	for {
		mt, msg, err = c.Conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		if c.ID == c.Pool.OwnerID {
			message := Message{Type: mt, Body: string(msg)}
			c.Pool.Broadcast <- message
			fmt.Printf("recv: %+v\n", message)
		}
	}
}
