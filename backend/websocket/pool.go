package customWebsocket

import "fmt"

type Pool struct {
	Register    chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			fmt.Println("Registering a user")
			pool.Clients[client] = true
			fmt.Println("Total connection pools :-", len(pool.Clients))

			for k, _ := range pool.Clients {
				fmt.Println(k)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New user joined now"})
			}
			break
		case client := <-pool.Unregister:
			fmt.Println("Unregistering a user")
			delete(pool.Clients, client)
			fmt.Println("Total connection just after unregister %d", pool.Clients)

			for k, _ := range pool.Clients {
				fmt.Println((k))
				client.Conn.WriteJSON(Message{Type: 1, Body: "One user has left the room"})
			}
			break
		case client := <-pool.Broadcast:
			fmt.Println("Broadcasting to users");
			for k,_:= range pool.Clients {
				if err := client.Conn.WriteJson(Message{Type: 1,Body:"Hello, this is a broadcasting message here"})
			}
			break
		}
		
	}
}
