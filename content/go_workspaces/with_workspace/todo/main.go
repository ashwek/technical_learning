package main

import (
	"github.com/ashwek/errorx"
	"github.com/ashwek/logger"
)

func main() {
	logger.Info("Hello world")
	err := errorx.NewError("This is an error")
	logger.Error(err)
}
