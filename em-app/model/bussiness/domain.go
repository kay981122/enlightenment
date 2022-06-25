package bussiness

import (
	"em-app/global"
)

type Domain struct {
	Id           int           `gorm:"column:id" form:"id" json:"id"`
	Domain       string        `gorm:"column:domain" form:"domain" json:"domain"`
	RegisterDate global.MyTime `gorm:"column:register_date" form:"register_date" json:"registerDate"`
	UpdateDate   global.MyTime `gorm:"column:update_date" form:"update_date" json:"updateDate"`
	Title        string        `gorm:"column:title" form:"title" json:"title"`
	Status       int8          `gorm:"column:status" form:"status" json:"status"`
}

// 自定义表名
func (Domain) TableName() string {
	return "domains"
}
