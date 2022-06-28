package common

import (
	"em-app/global"
	"em-app/model/common"
)

var (
	limit  = 10
	offset = 0
)

type DownloadCSVService struct {
}

func (downloadCSVService *DownloadCSVService) QueryDownloadCSV(info common.ExportCSVProgress) (list interface{}, err error) {
	db := global.Db.Model(&common.ExportCSVProgress{})
	var downloadCSV []common.ExportCSVProgress
	if info.Module != "" {
		db = db.Where("module = ?", info.Module)
	}
	if info.UserId != "" {
		db = db.Where("user_id = ?", info.UserId)
	}
	err = db.Limit(limit).Offset(offset).Order("CREATE_TIME desc").Find(&downloadCSV).Error
	return downloadCSV, err
}

func (downloadCSVService *DownloadCSVService) GetDownloadUrl(info common.ExportCSVProgress) (url string, err error) {
	var temp common.ExportCSVProgress
	db := global.Db.Model(&common.ExportCSVProgress{})
	err = db.Where("ID = ?", info.Id).First(&temp).Error
	return temp.FilePath, err
}
