package bussiness

import "em-app/service"

type ApiGroup struct {
	DomainApi
}

var (
	domainService      = service.ServiceGroupApp.BussinessServiceGroup.DomainService
	exportCSVService   = service.ServiceGroupApp.CommonServiceGroup.ExportCSVService
	downloadCSVService = service.ServiceGroupApp.CommonServiceGroup.DownloadCSVService
)
