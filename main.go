package main

import (
	"fmt"
	customwebsocket "gochat/websocket"
	"log"
	"net/http"
	"os"
)

func serverWs(pool *customwebsocket.Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("This is working")
	conn, err := customwebsocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	client := &customwebsocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	log.Println("This is working")
	pool := customwebsocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is Workinggg!!")
		serverWs(pool, w, r)
	})
}

func main() {
	setupRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	log.Println("Listening on port",port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port), nil))
}