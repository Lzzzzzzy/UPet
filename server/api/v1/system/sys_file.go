package system

import (
	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/common/response"
	"github.com/Lzzzzzzy/UPet/server/model/file"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileApi struct{}

// File Upload
// @Tags      File
// @Summary   文件上传
// @accept    application/json
// @Produce   application/json
// @Param     data  body      auth.AuthInfo            true  "微信签发的code"
// @Success   200   {object}  response.Response{msg=string}  "用户登录或注册"
// @Router    /api/file [post]
func (e *FileApi) UploadFile(c *gin.Context) {
	var file file.FileInfo
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = fileService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}

	fileUrl := map[string]string{"url": file.Url}
	response.OkWithDetailed(fileUrl, "上传成功", c)
}
