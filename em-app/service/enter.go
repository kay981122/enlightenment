package service

import (
	"em-app/service/bussiness"
	"em-app/service/common"
	"em-app/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	BussinessServiceGroup bussiness.ServiceGroup
	CommonServiceGroup    common.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
