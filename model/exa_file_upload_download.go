package model

import "github.com/cxpgo/golib/model"

type ExaFileUploadAndDownload struct {
	model.GormModel
	Name string `json:"name" gorm:"comment:文件名"`
	Url  string `json:"url" gorm:"comment:文件地址"`
	Tag  string `json:"tag" gorm:"comment:文件标签"`
	Key  string `json:"key" gorm:"comment:编号"`
}
