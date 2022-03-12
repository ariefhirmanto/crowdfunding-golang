package transaction

import "startup/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount      int    `json:"amount" binding:"required"`
	CampaignID  int    `json:"campaign_id" binding:"required"`
	PaymentType int    `json:"payment_type" binding:"required"`
	Message     string `json:"funder_message" binding:"required"`
	User        user.User
}

type CreateTransactionInputMVP struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	// User       user.User
	UserFullName string `json:"user_full_name" binding:"required"`
	UserEmail    string `json:"user_email" binding:"required"`
	PaymentType  int    `json:"payment_type" binding:"required"`
	Message      string `json:"message" binding:"required"`
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
