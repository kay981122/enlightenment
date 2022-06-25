package router

import (
	"em-app/router/bussiness"
	"em-app/router/system"
)

type RouterGroup struct {
	Bussiness bussiness.RouterGroup
	System    system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
