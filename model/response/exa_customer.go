package response

import "github.com/cxpgo/ginf/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
