package config

type Config struct {
	System  System  `mapstructure:"system"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha"`
	// oss
	Local Local `mapstructure:"local" json:"local"`
	Email Email `mapstructure:"email" json:"email"`
}
