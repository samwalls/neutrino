package handle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/googollee/go-socket.io"
)

// Handler handles all requried requests
type Handler struct {
	server *socketio.Server
	logger *logrus.Logger
}

type registerContent struct {
	username string
	files    map[string][]byte
}

type updateContent struct {
	files map[string][]byte
}

func (handler *Handler) deregister(so socketio.Socket) string {
	//TODO
	return ""
}

func (handler *Handler) register(so socketio.Socket, msg string) string {
	handler.logger.Info("register fired")
	content := registerContent{}
	err := json.Unmarshal([]byte(msg), &content)
	if err != nil {
		handler.logger.Fatal(err)
	}
	handler.logger.WithFields(logrus.Fields{
		"username": content.username,
		"files":    content.files,
	}).Info()
	response, err := json.Marshal(updateContent{
		files: content.files,
	})
	if err != nil {
		handler.logger.Fatal(err)
	}
	handler.logger.Info(response)
	return string(response)
}

func (handler *Handler) disconnection(so socketio.Socket) {
	// Handle disconnection
}

func (handler *Handler) connection(so socketio.Socket) {
	handler.logger.Info("connection established")
	handler.logger.Infof("Socket: %v", so)
	handler.logger.Infof("Connected clients: %v", handler.server.Count())
	so.On("register", func(msg string) string {
		so.Emit("polo", "HellO!")
		handler.logger.Info("register fired!!!")
	 	return handler.register(so, msg) 
	})
	so.On("deregister", func() string { return handler.deregister(so) })
	so.On("disconnection", handler.disconnection)
}

// NewHandler creates a new Handler struct
func NewHandler(logger *logrus.Logger) (Handler, error) {
	server, err := socketio.NewServer(nil)
	handler := Handler{
		server: server,
		logger: logger,
	}
	if err != nil {
		return handler, err
	}
	server.On("connection", func (so socketio.Socket) {
		handler.connection(so)
	})
	server.On("error", func(so socketio.Socket, err error) {
		handler.logger.Fatalf("error: %v", err)
	})
	return handler, nil
}

// Serve starts handling requests
func (handler *Handler) Serve(port string) {
	handler.logger.WithFields(logrus.Fields{
		"location": "localhost",
		"port":     port,
	}).Info()
	http.Handle("/socket.io/", handler.server)
	handler.logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
