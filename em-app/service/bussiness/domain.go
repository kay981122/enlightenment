package bussiness

import (
	"em-app/global"
	"em-app/model/bussiness"
	"em-app/model/bussiness/request"
	"gorm.io/gorm"
	"strconv"
)

type DomainService struct {
}

func (domainService *DomainService) GetDomainInfoList(info request.DomainSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.Db.Model(&bussiness.Domain{})
	var domainList []bussiness.Domain
	// 如果有条件搜索 下方会自动创建搜索语句
	SetDomainSearchData(info, db, "")
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("register_date desc").Find(&domainList).Error
	return domainList, total, err
}
func SetDomainSearchData(info request.DomainSearch, db *gorm.DB, fileName string) string {
	db = db.Select("id,domain,register_date,title,status")
	if info.Domain.Domain != "" {
		db = db.Where("domain LIKE ?", "%"+info.Domain.Domain+"%")
		fileName += "_" + info.Domain.Domain
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
		fileName += "_" + info.Title
	}
	if info.RegisterBeginTime != "" {
		db = db.Where("register_date >= ?", global.TimeStringToGoTime(info.RegisterBeginTime))
		fileName += "_" + info.RegisterBeginTime
	}
	if info.RegisterEndTime != "" {
		db = db.Where("register_date <= ?", global.TimeStringToGoTime(info.RegisterEndTime))
		fileName += "_" + info.RegisterEndTime
	}
	//if info.UpdateBeginTime != "" {
	//	db = db.Where("update_date >= ?", global.TimeStringToGoTime(info.UpdateBeginTime))
	//}
	//if info.UpdateEndTime != "" {
	//	db = db.Where("update_date <= ?", global.TimeStringToGoTime(info.UpdateEndTime))
	//}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
		fileName += "_" + strconv.FormatInt(info.Status, 10)
	}
	fileName += ".csv"
	return fileName
}
