package request

import "github.com/cxpgo/ginf/model"

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu
	AuthorityId string
}
