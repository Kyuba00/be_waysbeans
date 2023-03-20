package transactionsdto

type CreateTransactionRequest struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	UserID int    `json:"user_id"`
	Amount int    `json:"amount"`
}

type UpdateTransactionRequest struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	UserID int    `json:"user_id"`
	Amount int    `json:"amount"`
}
