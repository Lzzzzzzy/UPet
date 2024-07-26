package config

type Wechat struct {
	Miniprogram WechatMiniprogram `mapstructure:"miniprogram" json:"miniprogram" yaml:"miniprogram"` // 小程序配置
}

type WechatMiniprogram struct {
	AppId     string `mapstructure:"app-id" json:"app-id" yaml:"app-id"`             // 微信appid
	AppSecret string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"` // 微信app secret
}
