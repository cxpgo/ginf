package global

import (
	"github.com/cxpgo/ginf/model/config"
	"gorm.io/gorm"
	"time"
)

var (
	Config     = &config.Config{}
	ConfigName ="/config.toml"
	ConfigPath = "./config/"
)

type GormModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

}
