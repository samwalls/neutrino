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

func (handler *Handler) register(conn socketio.Conn, msg string) string {
	handler.logger.Info("register fired")
	content := registerContent{}
	err := json.Unmarshal([]byte(msg), &content)
	if err != nil {
		handler.logger.WithFields(logrus.Fields{
			"message": err.Error(),
		}).Fatal("error when unmarshalling message from register event")
	}
	handler.logger.WithFields(logrus.Fields{
		"username": content.username,
		"files":    content.files,
	}).Info()
	response, err := json.Marshal(updateContent{
		files: content.files,
	})
	if err != nil {
		handler.logger.WithFields(logrus.Fields{
			"message": err.Error(),
		}).Fatal("error when marshalling response to register event")
	}
	handler.logger.Info(response)
	return string(response)
}

func (handler *Handler) disconnection(conn socketio.Conn) {
	// Handle disconnection
}

func (handler *Handler) connection(conn socketio.Conn) error {
	handler.logger.Info("connection established")
	handler.logger.Infof("socket: %v", conn.RemoteAddr())
	return nil
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
	server.OnConnect("/", func(conn socketio.Conn) error {
		return handler.connection(conn)
	})
	server.OnEvent("/", "register", func(conn socketio.Conn, msg string) string {
		conn.Emit("polo", "Hello!")
		return handler.register(conn, msg)
	})
	server.OnError("/", func(err error) {
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
