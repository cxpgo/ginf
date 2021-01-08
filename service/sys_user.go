package service

import (

	"github.com/cxpgo/ginf/model"
	"github.com/cxpgo/ginf/utils"
	"github.com/cxpgo/golib/lib"
)

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = lib.GGorm.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}
