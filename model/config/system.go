package config

type System struct {
	Env            string `mapstructure:"env" json:"env"`
	Post           int    `mapstructure:"post" json:"post" `
	OssType        string `mapstructure:"oss-type" json:"ossType"`
	UseMultipoint  bool   `mapstructure:"use-multipoint" json:"useMultipoint"`
	UseMultidevice bool   `mapstructure:"use-multidevice" json:"useMultidevice"`
}
