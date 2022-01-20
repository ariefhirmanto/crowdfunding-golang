package payment

import (
	"startup/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
	GetPaymentURLMVP(transaction Transaction, user user.User) (string, error)
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	config, err := s.repository.GetPaymentConfig()
	if err != nil {
		return "", err
	}

	midclient.ClientKey = config.ClientKey
	midclient.ServerKey = config.ServerKey
	env := config.APIEnv
	if env == "production" {
		midclient.APIEnvType = midtrans.Production
	}
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}

func (s *service) GetPaymentURLMVP(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	config, err := s.repository.GetPaymentConfig()
	if err != nil {
		return "", err
	}

	midclient.ClientKey = config.ClientKey
	midclient.ServerKey = config.ServerKey
	env := config.APIEnv
	if env == "production" {
		midclient.APIEnvType = midtrans.Production
	}
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
