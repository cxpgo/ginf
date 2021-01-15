package model

import "github.com/cxpgo/golib/model"

type JwtBlacklist struct {
	model.GormModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
