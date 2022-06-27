package common

import (
	"em-app/global"
	"em-app/model/system"
	"fmt"
)

var (
	limit  = 10
	offset = 0
)

type DownloadCSVService struct {
}

func (downloadCSVService *DownloadCSVService) DownloadCSV(info system.ExportCSVProgress) (list interface{}, err error) {
	db := global.Db.Model(&system.ExportCSVProgress{})
	var downloadCSV []system.ExportCSVProgress
	if info.Module != "" {
		db = db.Where("module = ?", info.Module)
	}
	if info.UserId != "" {
		db = db.Where("user_id = ?", info.UserId)
	}
	err = db.Limit(limit).Offset(offset).Order("CREATE_TIME desc").Find(&downloadCSV).Error
	fmt.Println(downloadCSV)
	return downloadCSV, err
}
