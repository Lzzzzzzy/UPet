package response

import "github.com/Lzzzzzzy/UPet/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
