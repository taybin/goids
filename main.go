package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type appContext struct {
	area *Area
}

var lastID = 0
var upgrader = websocket.Upgrader{} // use default options

func main() {
	context := &appContext{
		area: NewArea(100, 100),
	}

	setupBoids(context)
	setupWeb(context)
}

func setupBoids(context *appContext) {
	for i := 0; i < 3; i++ {
		lastID = lastID + 1
		boid := NewBoid(lastID, len(context.area.Dimensions))
		boid.RandomizePosition(context.area)
		context.area.AddBoid(boid)
	}

	go func() {
		for {
			context.area.UpdateBoids()
			time.Sleep(time.Second)
		}
	}()
}

type appHandler struct {
	*appContext
	H func(*appContext, http.ResponseWriter, *http.Request)
}

func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ah.H(ah.appContext, w, r)
}

func setupWeb(context *appContext) {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.Handle("/ws", appHandler{context, websockHandler})

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func websockHandler(context *appContext, w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		select {
		case boid := <-context.area.SendChan:
			jBoid, err := json.Marshal(boid)
			if err != nil {
				log.Println("error encoding boid:", err)
				return
			}
			err = c.WriteMessage(websocket.TextMessage, jBoid)
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}
