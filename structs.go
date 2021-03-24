package alipay

import (
	"github.com/joaosoft/logger"
)

// AliPay ...
type AliPay struct {
	config        *AliPayConfig
	isLogExternal bool
	logger        logger.ILogger
}

// AppConfig ...
type AppConfig struct {
	AliPay *AliPayConfig `json:"alipay"`
}

// AliPayConfig ...
type AliPayConfig struct {
	Host         string `json:"host"`
	ReturnUrl    string `json:"return_url"`
	PublicApiKey string `json:"public_api_key"`
	SecretApiKey string `json:"secret_api_key"`
	Log          struct {
		Level string `json:"level"`
	} `json:"log"`
}

// AliPayOption ...
type AliPayOption func(auth *AliPay)
