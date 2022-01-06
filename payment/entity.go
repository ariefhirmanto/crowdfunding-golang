package payment

type Transaction struct {
	ID     int
	Amount int
}

type MidtransConfig struct {
	ClientKey string
	ServerKey string
	APIEnv    string
}
