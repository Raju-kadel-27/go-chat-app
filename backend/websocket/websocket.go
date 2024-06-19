package customWebsocket

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request)(*websocket.Conn, error){
	upgrader.checkOrigin= func (r *http.Request)bool{
		conn,err: upgrader.Upgrade(w,r,nil)
		fmt.Println(conn)
		if err != nil{
			fmt.Println("Websocket connection error %s",err)
			return;
		}
		return conn, nil
	}
}
