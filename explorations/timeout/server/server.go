package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	port := "5000"
	log.Println("Started HTTP server on port " + port)
	err := http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")
		time.Sleep(time.Second * 4)
		log.Println("Response sent")
		w.Write([]byte("This is the response!"))
	}))
	if err != nil {
		log.Fatalln("Error listening on port " + port)
	}
}
