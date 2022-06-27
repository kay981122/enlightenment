package bussiness

import (
	"em-app/model/bussiness/request"
	"em-app/model/common/response"
	"em-app/model/system"
	"github.com/gin-gonic/gin"
)

type DomainApi struct{}

func (d *DomainApi) GetDomainList(c *gin.Context) {
	var domainSearch request.DomainSearch
	// 解析JSON数据
	_ = c.ShouldBind(&domainSearch)
	// 验证入参数据

	if list, total, err := domainService.GetDomainInfoList(domainSearch); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     domainSearch.Page,
			PageSize: domainSearch.PageSize,
		}, "获取成功", c)
	}
}
func (d *DomainApi) ExportDomainCSV(c *gin.Context) {
	err := exportCSVService.GenerateCSV("domain")
	if err != nil {
		response.FailWithMessage("导出失败", c)
	} else {
		response.Ok(c)
	}
}
func (d *DomainApi) DownloadDomainCSV(c *gin.Context) {
	var downloadCSV system.ExportCSVProgress
	_ = c.ShouldBind(&downloadCSV)
	if list, err := downloadCSVService.DownloadCSV(downloadCSV); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
