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

func (handler *Handler) deregister(so socketio.Socket) {
	//TODO
}

func (handler *Handler) register(msg string, so socketio.Socket) {
	content := registerContent{}
	err := json.Unmarshal([]byte(msg), &content)
	if err != nil {
		handler.logger.Fatal(err)
	}
	handler.logger.WithFields(logrus.Fields{
		"username": content.username,
		"files":    content.files,
	})
	response, err := json.Marshal(updateContent{
		files: content.files,
	})
	if err != nil {
		handler.logger.Fatal(err)
	}
	handler.logger.Info(response)
	so.Emit("update", response)
}

// NewHandler creates a new Handler struct
func NewHandler(logger *logrus.Logger) (Handler, error) {
	handler := Handler{}
	server, err := socketio.NewServer(nil)
	if err != nil {
		return handler, err
	}
	server.On("register", handler.register)
	server.On("deregister", handler.deregister)
	return Handler{
		server: server,
		logger: logger,
	}, nil
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
