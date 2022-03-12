package transaction

import (
	"startup/campaign"
	"startup/user"
	"time"

	"github.com/leekchan/accounting"
)

type Transaction struct {
	ID          int
	CampaignID  int
	UserID      int
	Amount      int
	Status      string
	Code        string
	PaymentType int
	PaymentURL  string
	Message     string
	User        user.User
	Campaign    campaign.Campaign
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TransactionMVP struct {
	ID           int
	CampaignID   int
	Amount       int
	Status       string
	Code         string
	PaymentType  int
	PaymentURL   string
	Message      string
	UserFullName string
	UserEmail    string
	Campaign     campaign.Campaign
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

const (
	PaymentBankTransfer int = 1
	PaymentGopay        int = 2
	PaymentShopeePay    int = 3
)

var mapPayment = map[int]string{
	PaymentBankTransfer: "bank_transfer",
	PaymentGopay:        "gopay",
	PaymentShopeePay:    "shopee_pay",
}

func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
