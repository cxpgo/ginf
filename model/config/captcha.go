package config

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight"`
}
