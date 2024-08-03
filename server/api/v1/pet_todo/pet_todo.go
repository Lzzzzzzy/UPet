package petTodo

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/common"
	"github.com/Lzzzzzzy/UPet/server/model/common/response"
	petTodoModel "github.com/Lzzzzzzy/UPet/server/model/pet_todo"
	petTodoReq "github.com/Lzzzzzzy/UPet/server/model/pet_todo/request"
	petTodoResp "github.com/Lzzzzzzy/UPet/server/model/pet_todo/response"
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
// @Router    /api/pet-todo [post]
func (e *PetTodoApi) CreatePetTodo(c *gin.Context) {
	var petTodoResp petTodoReq.PetTodoInfo
	err := c.ShouldBindJSON(&petTodoResp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(petTodoResp, utils.PetTodoInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var petTodo petTodoModel.PetTodoInfo
	petTodo.Title = petTodoResp.Title
	petTodo.Remark = petTodoResp.Remark
	petTodo.Remind = petTodoResp.Remind
	petTodo.Complete = petTodoResp.Complete
	petTodo.Type = petTodoResp.Type
	petTodo.Color = petTodoResp.Color
	petTodo.TodoTime = *petTodoResp.TodoTime
	petTodo.PetId = petTodoResp.PetId
	petTodo.CreatedBy = utils.GetUserID(c)
	petTodo.UpdatedBy = utils.GetUserID(c)
	petTodo.FamilyId = utils.GetUserFamilyID(c)

	needCreateTodos := []*petTodoModel.PetTodoInfo{}
	if petTodoResp.Remind {
		remindTimes := petTodoService.FormatRemindTime(&petTodoResp)
		for _, remindTime := range remindTimes {
			newTodo := petTodo
			newTodo.RemindTime = &common.CustomTime{Time: remindTime}
			needCreateTodos = append(needCreateTodos, &newTodo)
		}
	} else {
		needCreateTodos = append(needCreateTodos, &petTodo)
	}

	err = petTodoService.CreatePetTodos(needCreateTodos)

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
// @Param     petTodoID  path      int             true  "宠物待办ID"
// @Success   200   {object}  response.Response{msg=string}  "删除宠物待办"
// @Router    /api/pet-todo/:petTodoID [delete]
func (e *PetTodoApi) DeletePetTodo(c *gin.Context) {
	var petTodo petTodoModel.PetTodoInfo
	todoIdStr := c.Param("petTodoID")
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
// @Param     petTodoID  path      int             true  "宠物待办ID"
// @Param     data  body      petTodo.PetTodoInfo            true  "宠物待办信息"
// @Success   200   {object}  response.Response{msg=string}  "更新宠物待办信息"
// @Router    /api/pet-todo/:petTodoID [put]
func (e *PetTodoApi) UpdatePetTodo(c *gin.Context) {
	var petTodo petTodoModel.PetTodoInfo
	todoIdStr := c.Param("petTodoID")
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
	olderPetTodo, err := petTodoService.GetPetTodoInfo(petTodo.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	petTodo.CreatedAt = olderPetTodo.CreatedAt
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
// @Param     petTodoID  path      int             true  "宠物待办ID"
// @Success   200   {object}  response.Response{data=petRes.PetTodoResponse,msg=string}  "获取单一宠物信息"
// @Router    /api/pet-todo/:petTodoID [get]
func (e *PetTodoApi) GetPetTodo(c *gin.Context) {
	todoIdStr := c.Param("petTodoID")
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
	response.OkWithDetailed(petTodoResp.PetTodoInfoResponse{PetTodo: data}, "获取成功", c)
}

// GetPetTodoList
// @Tags      PetTodo
// @Summary   获取所有宠物待办列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /api/pet-todos [get]
func (e *PetTodoApi) GetPetTodoList(c *gin.Context) {
	var pageInfo petTodoReq.PetTodoCondition
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
	minTime, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 00:00:00", pageInfo.Date))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	maxTime, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 23:59:59", pageInfo.Date))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	petTodoList, err := petTodoService.GetPetTodoInfoList(minTime, maxTime, pageInfo.PetId)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(petTodoList, "获取成功", c)
}

// GetPetTodoMark
// @Tags      PetTodo
// @Summary   获取指定日期中宠物待办标记
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /api/pet-todos/mark [post]
func (e *PetTodoApi) GetPetTodoMarkList(c *gin.Context) {
	var dates petTodoReq.PetTodoMarkCondition
	err := c.ShouldBindJSON(&dates)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	minTime, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 00:00:00", dates.MinDate))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	maxTime, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 23:59:59", dates.MaxDate))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	petTodoColors, err := petTodoService.GetPetTodoInfoMarkList(minTime, maxTime, utils.GetUserFamilyID(c))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	markList := map[string]int{}
	for _, info := range petTodoColors {
		todoDate := info.TodoTime.Format("2006-01-02")
		if int(info.Color) >= markList[todoDate] {
			markList[todoDate] = int(info.Color)
		}
	}
	response.OkWithDetailed(markList, "获取成功", c)
}
