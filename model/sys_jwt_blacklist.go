package model

import (
	"github.com/cxpgo/ginf/global"
)

type JwtBlacklist struct {
	global.GormModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
