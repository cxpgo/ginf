package config

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" `
	ExpiresTime int64 `mapstructure:"expires_time" json:"expiresTime" `
	BufferTime int64 `mapstructure:"buffer_time" json:"bufferTime" `
}
