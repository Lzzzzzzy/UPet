package system

import (
	"github.com/Lzzzzzzy/UPet/server/middleware"
	"github.com/gin-gonic/gin"
)

type SystemRouter struct{}

func (s *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysConfigRouterWithoutRecord := Router.Group("system").Use(middleware.OperationRecord())
	{
		sysConfigRouterWithoutRecord.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
		sysConfigRouterWithoutRecord.POST("getServerInfo", systemApi.GetServerInfo)     // 获取服务器信息
		sysConfigRouterWithoutRecord.POST("reloadSystem", systemApi.ReloadSystem)       // 重启系统
		sysConfigRouterWithoutRecord.POST("setSystemConfig", systemApi.SetSystemConfig) // 设置配置文件内容
	}
}
