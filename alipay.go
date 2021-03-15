package alipay

import (
	"alipay/routes"
	"fmt"
	"github.com/joaosoft/logger"
	"github.com/stripe/stripe-go"
	"net/http"
	"strings"
)

func (ap *AliPay) Start() error {
	routes.InitializeRouters(ap.config.PublicApiKey)

	url := fmt.Sprintf("%s/checkout", ap.config.Host)
	fmt.Printf("\n:: waiting for feedback at %s", url)

	var err error
	if err = showUrl(url); err != nil {
		return err
	}

	split := strings.Split(ap.config.Host, ":")
	if len(split) == 0 {
		return fmt.Errorf("invalid port")
	}
	if err = http.ListenAndServe(fmt.Sprintf(":%s", split[len(split)-1]), nil); err != nil {
		return err
	}

	return nil
}

// NewAliPay ...
func NewAliPay(options ...AliPayOption) (*AliPay, error) {
	config, _, err := NewConfig()

	service := &AliPay{
		logger: logger.NewLogDefault("alipay", logger.WarnLevel),
		config: config.AliPay,
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.AliPay != nil {
		level, _ := logger.ParseLevel(config.AliPay.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	} else {
		config.AliPay = &AliPayConfig{
			Host: defaultURL,
		}
	}

	service.Reconfigure(options...)

	stripe.Key = config.AliPay.SecretApiKey

	return service, nil
}
