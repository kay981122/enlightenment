package common

import "em-app/model/bussiness/request"

type ExportCSVSearch struct {
	UserId string `gorm:"column:USER_ID" form:"userId" json:"userId"`
	Module string `gorm:"column:MODULE" form:"module" json:"module"`
	request.DomainSearch
}
