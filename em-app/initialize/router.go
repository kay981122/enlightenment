package initialize

import (
	"em-app/router"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.Default()
	bussinessRouter := router.RouterGroupApp.Bussiness
	//systemRouter := router.RouterGroupApp.System
	PublicGroup := Router.Group("")
	{
		// 健康检查
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		bussinessRouter.InitDomainRouter(PublicGroup)
		//systemRouter.Init
	}
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//{
	//	systemRouter.InitApiRouter(PrivateGroup)                 // 注册功能api路由
	//	systemRouter.InitJwtRouter(PrivateGroup)                 // jwt相关路由
	//	systemRouter.InitUserRouter(PrivateGroup)                // 注册用户路由
	//	systemRouter.InitMenuRouter(PrivateGroup)                // 注册menu路由
	//	systemRouter.InitSystemRouter(PrivateGroup)              // system相关路由
	//	systemRouter.InitCasbinRouter(PrivateGroup)              // 权限相关路由
	//	systemRouter.InitAutoCodeRouter(PrivateGroup)            // 创建自动化代码
	//	systemRouter.InitAuthorityRouter(PrivateGroup)           // 注册角色路由
	//	systemRouter.InitSysDictionaryRouter(PrivateGroup)       // 字典管理
	//	systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)     // 自动化代码历史
	//	systemRouter.InitSysOperationRecordRouter(PrivateGroup)  // 操作记录
	//	systemRouter.InitSysDictionaryDetailRouter(PrivateGroup) // 字典详情管理
	//	systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)  // 字典详情管理
	//
	//	exampleRouter.InitExcelRouter(PrivateGroup)                 // 表格导入导出
	//	exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
	//	exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
	//
	//	// Code generated by github.com/flipped-aurora/gin-vue-admin/server Begin; DO NOT EDIT.
	//
	//	// Code generated by github.com/flipped-aurora/gin-vue-admin/server End; DO NOT EDIT.
	//}
	return Router
}