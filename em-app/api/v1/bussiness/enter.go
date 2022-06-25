package bussiness

import "em-app/service"

type ApiGroup struct {
	DomainApi
}

var (
	domainService = service.ServiceGroupApp.BussinessServiceGroup.DomainService
)
