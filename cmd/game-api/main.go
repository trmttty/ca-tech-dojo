package main

import (
	"github.com/trmttty/ca-tech-dojo/handler"
	"github.com/trmttty/ca-tech-dojo/injector"
)

func main() {

	userHandler := injector.InjectUserHandler()

	handler.Run(userHandler)
}
