package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/samwalls/neutrino/handle"
)

func main() {
	logger := logrus.New()
	handler := handle.NewHandler(logger)
	handler.Serve("/", 8000)
}
