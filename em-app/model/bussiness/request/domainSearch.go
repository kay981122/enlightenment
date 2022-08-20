package request

import (
	"em-app/model/bussiness"
	"em-app/model/common/request"
)

type DomainSearch struct {
	bussiness.Domain
	RegisterBeginTime string `json:"registerBeginTime" form:"registerBeginTime"`
	RegisterEndTime   string `json:"registerEndTime" form:"registerEndTime"`
	//UpdateBeginTime   string `json:"updateBeginTime" form:"updateBeginTime"`
	//UpdateEndTime     string `json:"updateEndTime" form:"updateEndTime"`
	request.PageInfo
}
