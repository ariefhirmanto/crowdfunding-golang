package payment

import (
	"math"
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

	grossAmount := int64(transaction.Amount) + int64(math.Ceil(0.3*float64(transaction.Amount)))
	if transaction.PaymentType == "bank_transfer" {
		grossAmount = int64(transaction.Amount) + int64(5000)
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		EnabledPayments: []midtrans.PaymentType{
			midtrans.PaymentType(transaction.PaymentType),
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: grossAmount,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
