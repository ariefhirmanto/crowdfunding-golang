package payment

type Transaction struct {
	ID          int
	Amount      int
	PaymentType string
}

type MidtransConfig struct {
	ClientKey string
	ServerKey string
	APIEnv    string
}
