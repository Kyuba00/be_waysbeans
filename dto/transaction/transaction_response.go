package transactionsdto

type TransactionResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	UserID int    `json:"user_id"`
	Amount int    `json:"amount"`
}
