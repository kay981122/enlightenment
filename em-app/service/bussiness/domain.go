package bussiness

import (
	"em-app/global"
	"em-app/model/bussiness"
	"em-app/model/bussiness/request"
)

type DomainService struct {
}

func (domainService *DomainService) GetDomainInfoList(info request.DomainSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.Db.Model(&bussiness.Domain{})
	var domainList []bussiness.Domain
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Domain.Domain != "" {
		db = db.Where("domain LIKE ?", "%"+info.Domain.Domain+"%")
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.RegisterBeginTime != "" {
		db = db.Where("register_date >= ?", global.TimeStringToGoTime(info.RegisterBeginTime))
	}
	if info.RegisterEndTime != "" {
		db = db.Where("register_date <= ?", global.TimeStringToGoTime(info.RegisterEndTime))
	}
	if info.UpdateBeginTime != "" {
		db = db.Where("update_date >= ?", global.TimeStringToGoTime(info.UpdateBeginTime))
	}
	if info.UpdateEndTime != "" {
		db = db.Where("update_date <= ?", global.TimeStringToGoTime(info.UpdateEndTime))
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("register_date desc").Find(&domainList).Error
	return domainList, total, err
}
