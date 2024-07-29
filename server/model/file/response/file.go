package response

import "github.com/Lzzzzzzy/UPet/server/model/file"

type FileResponse struct {
	File file.FileInfo `json:"file"`
}
