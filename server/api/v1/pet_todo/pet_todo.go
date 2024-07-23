package petTodo

import (
	"strconv"

	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/common/request"
	"github.com/Lzzzzzzy/UPet/server/model/common/response"
	petTodo "github.com/Lzzzzzzy/UPet/server/model/pet_todo"
	petTodoRes "github.com/Lzzzzzzy/UPet/server/model/pet_todo/response"
	"github.com/Lzzzzzzy/UPet/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PetTodoApi struct{}

// CreatePetTodo
// @Tags      PetTodo
// @Summary   创建宠物待办
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      pet.PetTodo            true  "宠物待办信息"
// @Success   200   {object}  response.Response{msg=string}  "创建宠物待办"
// @Router    /pet/todo [post]
func (e *PetTodoApi) CreatePetTodo(c *gin.Context) {
	var petTodo petTodo.PetTodoInfo
	err := c.ShouldBindJSON(&petTodo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(petTodo, utils.PetTodoInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	petTodo.CreatedBy = utils.GetUserID(c)
	petTodo.UpdatedBy = utils.GetUserID(c)
	err = petTodoService.CreatePetTodo(petTodo)

	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePetTodo
// @Tags      PetTodo
// @Summary   删除宠物待办
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     todoID  path      int             true  "宠物待办ID"
// @Success   200   {object}  response.Response{msg=string}  "删除宠物待办"
// @Router    /pet/todo/:todoID [delete]
func (e *PetTodoApi) DeletePetTodo(c *gin.Context) {
	var petTodo petTodo.PetTodoInfo
	todoIdStr := c.Param("todoID")
	petTodoID, err := strconv.ParseUint(todoIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("宠物待办ID错误", c)
		return
	}
	petTodo.ID = uint(petTodoID)
	err = utils.Verify(petTodo.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = petTodoService.DeletePetTodo(petTodo)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdatePetTodo
// @Tags      PetTodo
// @Summary   更新宠物待办信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     todoID  path      int             true  "宠物待办ID"
// @Param     data  body      petTodo.PetTodoInfo            true  "宠物待办信息"
// @Success   200   {object}  response.Response{msg=string}  "更新宠物待办信息"
// @Router    /pet/todo/:todoID [put]
func (e *PetTodoApi) UpdatePetTodo(c *gin.Context) {
	var petTodo petTodo.PetTodoInfo
	todoIdStr := c.Param("todoID")
	petTodoID, err := strconv.ParseUint(todoIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("宠物ID错误", c)
		return
	}

	err = c.ShouldBindJSON(&petTodo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	petTodo.ID = uint(petTodoID)
	err = utils.Verify(petTodo.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(petTodo, utils.PetTodoInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = petTodoService.UpdatePetTodo(&petTodo)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetPetTodo
// @Tags      PetTodo
// @Summary   获取单一宠物待办信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     pet.PetTodo                                             true  "宠物待办ID"
// @Success   200   {object}  response.Response{data=petRes.PetTodoResponse,msg=string}  "获取单一宠物信息"
// @Router    /pet/todo/:todoID [get]
func (e *PetTodoApi) GetPetTodo(c *gin.Context) {
	todoIdStr := c.Param("todoID")
	petTodoID, err := strconv.ParseUint(todoIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("宠物待办ID错误", c)
		return
	}

	data, err := petTodoService.GetPetTodoInfo(uint(petTodoID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(petTodoRes.PetTodoInfoResponse{PetTodo: data}, "获取成功", c)
}

// GetPetTodoList
// @Tags      PetTodo
// @Summary   获取所有宠物待办列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /pets [get]
func (e *PetTodoApi) GetPetTodoList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	petList, total, err := petTodoService.GetPetInfoList(pageInfo.Keyword, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     petList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
