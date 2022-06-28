package router

import (
	"em-app/router/bussiness"
	"em-app/router/commom"
	"em-app/router/system"
)

type RouterGroup struct {
	Bussiness bussiness.RouterGroup
	Common    commom.RouterGroup
	System    system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
