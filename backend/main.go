package main

import (
	customWebsocket "chatapp/websocket"
	"fmt"
	"net/http"
)

func serveWs(pool *customWebsocket.Pool, w http.ResponseWriter, r *http.Request) {

	fmt.Println("This websocket server is working properly...")

	conn, err := customWebsocket.Upgrade(w, r)

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &customWebsocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {

	fmt.Println("Setup routes is called() ")

	pool := customWebsocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

}

func main() {
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
