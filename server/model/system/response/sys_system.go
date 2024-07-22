package response

import "github.com/Lzzzzzzy/UPet/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
