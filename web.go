package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{} // use default options

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
			pBoid, err := proto.Marshal(boid.ToBoidPosition())
			if err != nil {
				log.Println("error encoding boid:", err)
				return
			}
			err = c.WriteMessage(websocket.BinaryMessage, pBoid)
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}
