package core

import (
	"em-app/initialize"
)

type server interface {
}

func RunWindowsServer() {
	Router := initialize.Routers()
	Router.Run(":9080")
}
