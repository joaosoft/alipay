package alipay

import (
	"github.com/joaosoft/logger"
)

// Reconfigure ...
func (ap *AliPay) Reconfigure(options ...AliPayOption) {
	for _, option := range options {
		option(ap)
	}
}

// WithConfiguration ...
func WithConfiguration(config *AliPayConfig) AliPayOption {
	return func(ap *AliPay) {
		ap.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) AliPayOption {
	return func(ap *AliPay) {
		ap.logger = logger
		ap.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) AliPayOption {
	return func(ap *AliPay) {
		ap.logger.SetLevel(level)
	}
}
