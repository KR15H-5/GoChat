package main

import (
	"fmt"
	"net/http"

	"github.com/KRI5H-5/GoChat/pkg/websocket"
)

func SetupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}
func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint reached")

	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func main() {

	fmt.Println("Krish's Full stack chat project")
	SetupRoutes()
	http.ListenAndServe(":9000", nil)

}
