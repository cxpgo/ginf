package config

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" `
}
