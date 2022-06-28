package common

import (
	"em-app/model/common"
	"em-app/model/common/response"
	"github.com/gin-gonic/gin"
)

type CommonApi struct {
}

func (commonApi *CommonApi) ExportCSV(c *gin.Context) {
	var exportCSVSearch common.ExportCSVSearch
	_ = c.ShouldBind(&exportCSVSearch)
	err := exportCSVService.GenerateCSV(exportCSVSearch)
	if err != nil {
		response.FailWithMessage("导出失败", c)
	} else {
		response.Ok(c)
	}
}
func (commonApi *CommonApi) QueryExportCSV(c *gin.Context) {
	var downloadCSV common.ExportCSVProgress
	_ = c.ShouldBind(&downloadCSV)
	if list, err := downloadCSVService.QueryDownloadCSV(downloadCSV); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
func (commonApi *CommonApi) DownloadCSV(c *gin.Context) {
	var downloadCSV common.ExportCSVProgress
	_ = c.ShouldBind(&downloadCSV)
	url, _ := downloadCSVService.GetDownloadUrl(downloadCSV)
	c.File(url)
}
