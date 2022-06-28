package common

import "em-app/service"

type ApiGroup struct {
	CommonApi
}

var (
	downloadCSVService = service.ServiceGroupApp.CommonServiceGroup.DownloadCSVService
	exportCSVService   = service.ServiceGroupApp.CommonServiceGroup.ExportCSVService
)
