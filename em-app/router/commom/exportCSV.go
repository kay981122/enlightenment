package commom

import (
	v1 "em-app/api/v1"
	"github.com/gin-gonic/gin"
)

type ExportCSVRouter struct {
}

func (s *ExportCSVRouter) InitExportCSVRouter(Router *gin.RouterGroup) {

	exportRouterWithoutRecord := Router.Group("csv")
	commmonApi := v1.ApiGroupApp.CommonApiGroup.CommonApi
	{
		exportRouterWithoutRecord.POST("export", commmonApi.ExportCSV)                 // 导出CSV
		exportRouterWithoutRecord.POST("queryExportResult", commmonApi.QueryExportCSV) // 查询导出结果
		exportRouterWithoutRecord.POST("download", commmonApi.DownloadCSV)             // 下载CSV
	}
}
