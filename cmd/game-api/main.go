package main

import (
	"github.com/trmttty/ca-tech-dojo/pkg/injector"
	"github.com/trmttty/ca-tech-dojo/pkg/interface/handler"
)

func main() {

	userHandler := injector.InjectUserHandler()

	handler.Run(userHandler)
}
