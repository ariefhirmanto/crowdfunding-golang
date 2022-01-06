package payment

import (
	"errors"
	"startup/config"
)

type Repository interface {
	GetPaymentConfig() (MidtransConfig, error)
}

type repository struct {
	config config.PaymentConfigurations
}

func NewRepository(config config.PaymentConfigurations) *repository {
	return &repository{config}
}

func (r *repository) GetPaymentConfig() (MidtransConfig, error) {
	var config MidtransConfig

	config.APIEnv = r.config.APIEnv
	config.ClientKey = r.config.ClientKey
	config.ServerKey = r.config.ServerKey
	if (config == MidtransConfig{}) {
		return config, errors.New("failed to get config")
	}

	return config, nil
}
