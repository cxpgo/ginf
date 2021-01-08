package config

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path" `
}

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone"`
	Bucket        string `mapstructure:"bucket" json:"bucket"`
	ImgPath       string `mapstructure:"img-path" json:"imgPath"`
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps"`
	AccessKey     string `mapstructure:"access-key" json:"accessKey"`
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" `
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains"`
}
