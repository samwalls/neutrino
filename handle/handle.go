package handle

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

// Handler handles all requried requests
type Handler struct {
	server *gosocketio.Server
	logger *logrus.Logger
}

func (handler *Handler) register(c *gosocketio.Channel, s struct{}) string {
	handler.logger.WithFields(logrus.Fields{
		"ID": c.Id(),
	}).Info("register")
	//handler.logger.WithFields(logrus.Fields{
	//	"username": session.User.,
	//	"files":    session.,
	//}).Info("received register content")
	//if err != nil {
	//	handler.logger.WithFields(logrus.Fields{
	//		"message": err.Error(),
	//	}).Fatal("error when marshalling response to register event")
	//}
	handler.logger.Infof("got: %v", s)
	return string("I got your register")
}

func (handler *Handler) disconnection(c *gosocketio.Channel) {
	//TODO handle disconnect

}

func (handler *Handler) connection(c *gosocketio.Channel) error {
	handler.logger.WithFields(logrus.Fields{
		"ID": c.Id(),
	}).Info("connection established")
	c.Emit("polo", "hey m88")
	return nil
}

// NewHandler creates a new Handler struct
func NewHandler(logger *logrus.Logger) (Handler, error) {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	handler := Handler{
		server: server,
		logger: logger,
	}
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) error {
		return handler.connection(c)
	})
	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		handler.disconnection(c)
	})
	server.On("/register", func(c *gosocketio.Channel, s struct{}) string {
		return handler.register(c, s)
	})
	server.On(gosocketio.OnError, func(err error) {
		handler.logger.Fatalf("error: %v", err)
	})
	return handler, nil
}

// Serve starts handling requests
func (handler *Handler) Serve(port string) {
	handler.logger.WithFields(logrus.Fields{
		"location": "localhost",
		"port":     port,
	}).Info("server started")
	http.Handle("/socket.io/", handler.server)
	handler.logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
