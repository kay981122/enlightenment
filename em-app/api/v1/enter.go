package v1

import (
	"em-app/api/v1/bussiness"
	"em-app/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	BussinessApiGroup bussiness.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
