package core

import (
	"em-app/initialize"
)

type server interface {
}

func RunWindowsServer() {
	Router := initialize.Routers()
	Router.Run(":9080")
	//h := &http.Server{
	//	Handler:        Router,
	//	ReadTimeout:    20 * time.Second,
	//	WriteTimeout:   20 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
}
