package initialize

import (
	"fmt"

	"github.com/Lzzzzzzy/UPet/server/utils/plugin"

	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		fmt.Println(Plugin[i].RouterPath(), "注册开始!")
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
		fmt.Println(Plugin[i].RouterPath(), "注册成功!")
	}
}

func bizPluginV1(group ...*gin.RouterGroup) {
	private := group[0]
	public := group[1]
	holder(public, private)
}
