package config

type Jwt struct {
	SigningKey string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"` // jwt签名
	ExpireTime int64  `mapstructure:"expire-time" json:"expire-time" yaml:"expire-time"` //过期时间单位为ms
	BufferTime int64  `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"` //缓存时间
	Issuer     string `mapstructure:"issuer"      json:"issuer"      yaml:"issuer"`      //签发者
}
