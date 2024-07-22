package request

import (
	"time"

	"github.com/Lzzzzzzy/UPet/server/model/common/request"
)

type InfoSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
