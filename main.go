package main

import (
	"flag"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	ws "github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8000", "http service address")

func serve(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"client": r.Host,
	}).Info("client connected")
}

func main() {
	fmt.Println(ws.BinaryMessage)
	http.HandleFunc("/", serve)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.SetLevel(log.FatalLevel)
		log.Fatalln("error accepting client")
	}
}
