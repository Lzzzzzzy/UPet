package response

type PetInfoResponse struct {
	ID              uint   `json:"id" form:"id" `                           // 宠物头像
	Avatar          string `json:"avatar" form:"avatar" `                   // 宠物头像
	Name            string `json:"name" form:"name" `                       // 宠物名
	Birthday        string `json:"birthday" form:"birthday" `               // 宠物生日
	Gender          uint   `json:"gender" form:"gender" `                   // 宠物性别
	Category        uint   `json:"category" form:"category" `               // 宠物类型
	SterilizedState uint   `json:"sterilizedState" form:"sterilizedState" ` // 宠物绝育状态
}
