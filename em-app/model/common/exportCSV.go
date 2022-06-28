package common

type ExportCSVProgress struct {
	Id         string `gorm:"column:ID" form:"id" json:"id"`
	UserId     string `gorm:"column:USER_ID" form:"userId" json:"userId"`
	FilePath   string `gorm:"column:FILE_PATH" form:"filePath" json:"filePath"`
	Module     string `gorm:"column:MODULE" form:"module" json:"module"`
	Status     string `gorm:"column:STATUS" form:"status" json:"status"`
	CreateTime string `gorm:"column:CREATE_TIME" form:"createTime" json:"createTime"`
	UpdateTime string `gorm:"column:UPDATE_TIME" form:"updateTime" json:"updateTime"`
}

func (ExportCSVProgress) TableName() string {
	return "export_csv_progress"
}
