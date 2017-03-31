package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	setupBoids()
	setupWeb()
}

func setupWeb() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

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
