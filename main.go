package main

import (
	// "github.com/Sirupsen/logrus"
	// "github.com/samwalls/neutrino/handle"

    "log"
    "net/http"

    socketio "github.com/googollee/go-socket.io"
)

// func main() {
// 	logger := logrus.New()
// 	handler, err := handle.NewHandler(logger)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	handler.Serve("8000")
// }


func main() {
    server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
        log.Println("on connection")
        so.Join("chat")
        so.On("register", func(msg string) {
            log.Println("emit:", so.Emit("polo", msg))
            so.BroadcastTo("chat", "chat message", msg)
        })
        so.On("disconnection", func() {
            log.Println("on disconnect")
        })
    })
    server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })

    http.Handle("/socket.io/", server)
    http.Handle("/", http.FileServer(http.Dir("./asset")))
    log.Println("Serving at localhost:8000...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}