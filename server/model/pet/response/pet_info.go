package response

import (
	"github.com/Lzzzzzzy/UPet/server/model/pet"
)

type PetInfoResponse struct {
	Pet pet.PetInfo `json:"pet"`
}
