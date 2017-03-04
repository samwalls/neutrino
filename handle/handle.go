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
	so.Emit("update", response)
}

func (handler *Handler) connection(so socketio.Socket) {
	handler.logger.Info("connection established")
	so.On("register", handler.register)
	so.On("deregister", handler.deregister)
}

// NewHandler creates a new Handler struct
func NewHandler(logger *logrus.Logger) (Handler, error) {
	handler := Handler{}
	server, err := socketio.NewServer(nil)
	if err != nil {
		return handler, err
	}
	server.On("connection", handler.connection)
	server.On("error", func(so socketio.Socket, err error) {
		handler.logger.Fatal(err)
	})
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
