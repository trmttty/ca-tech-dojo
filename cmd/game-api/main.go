package main

import (
	"github.com/trmttty/ca-tech-dojo/injector"
	"github.com/trmttty/ca-tech-dojo/interface/handler"
)

func main() {

	userHandler := injector.InjectUserHandler()

	handler.Run(userHandler)
}
