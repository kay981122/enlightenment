package bussiness

import (
	v1 "em-app/api/v1"
	"github.com/gin-gonic/gin"
)

type DomainRouter struct {
}

func (s *DomainRouter) InitDomainRouter(Router *gin.RouterGroup) {

	domainRouterWithoutRecord := Router.Group("domain")
	domainApi := v1.ApiGroupApp.BussinessApiGroup.DomainApi
	{
		domainRouterWithoutRecord.POST("getDomainList", domainApi.GetDomainList) // 分页获取域名列表
	}
}
