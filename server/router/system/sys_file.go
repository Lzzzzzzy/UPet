package system

import (
	"github.com/Lzzzzzzy/UPet/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileRouter struct{}

func (r *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file").Use(middleware.OperationRecord())
	{
		fileRouter.POST("", fileApi.UploadFile) // 上传文件
	}
}
