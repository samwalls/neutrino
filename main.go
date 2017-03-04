package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/samwalls/neutrino/handle"
)

func main() {
	logger := logrus.New()
	handler, err := handle.NewHandler(logger)
	if err != nil {
		logger.Fatal(err)
	}
	handler.Serve("/", 8000)
}
