package main

import (
	"errorx"
	"logger"
)

func main() {
	err := errorx.NewError("This is an error")
	logger.Error(err)
}
