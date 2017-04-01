package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{} // use default options

func main() {
	setupBoids()
	setupWeb()
}

func setupWeb() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", websock)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func setupBoids() {
	a := NewArea(1000, 1000)
	for i := 0; i < 20; i++ {
		a.AddBoid()
	}
	log.Printf("%v\n", a)

	go func() {
		for {
			a.UpdateBoids()
			time.Sleep(time.Second)
		}
	}()
}

func websock(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
