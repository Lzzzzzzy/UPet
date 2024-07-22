package response

import "github.com/Lzzzzzzy/UPet/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
