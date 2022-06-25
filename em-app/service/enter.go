package service

import (
	"em-app/service/bussiness"
	"em-app/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	BussinessServiceGroup bussiness.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
