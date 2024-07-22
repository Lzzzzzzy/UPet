package response

import (
	"github.com/Lzzzzzzy/UPet/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
