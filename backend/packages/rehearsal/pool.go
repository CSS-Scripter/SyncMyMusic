package rehearsal

import "fmt"

// Pool struct for a pool
type Pool struct {
	Register    chan *Client
	Unregister  chan *Client
	Clients     map[*Client]bool
	Broadcast   chan Message
	RehearsalID int
	OwnerID     int
}

// NewPool constructor for Pool
func NewPool(rehearsalID int, ownerID int) *Pool {
	return &Pool{
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Clients:     make(map[*Client]bool),
		Broadcast:   make(chan Message),
		RehearsalID: rehearsalID,
		OwnerID:     ownerID,
	}
}

// Start starts a pool
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined...", RehearsalID: pool.RehearsalID, OwnerID: pool.OwnerID})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			if client.ID == pool.OwnerID {
				for client := range pool.Clients {
					client.Conn.WriteJSON(Message{Type: -1, Body: "Session ended...", RehearsalID: pool.RehearsalID, OwnerID: pool.OwnerID})
				}
				DeleteRehearsal(pool.RehearsalID)
				break
			}
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 0, Body: "User Disconnected...", RehearsalID: pool.RehearsalID, OwnerID: pool.OwnerID})
			}
			if clientMapIsEmpty(pool) {
				fmt.Println("Closing this rehearsal! ", pool.RehearsalID)
				DeleteRehearsal(pool.RehearsalID)
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(Message{Type: 2, Body: message.Body, RehearsalID: pool.RehearsalID, OwnerID: pool.OwnerID}); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

func clientMapIsEmpty(pool *Pool) bool {
	return len(pool.Clients) == 0
}
