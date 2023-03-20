package cartsdto

type CreateCartRequest struct {
	ProductID     int `json:"product_id"`
	TransactionID int `json:"transaction_id"`
	Qty           int `json:"qty"`
	SubAmount     int `json:"sub_amount"`
}

type UpdateCartRequest struct {
	ProductID     int `json:"product_id"`
	TransactionID int `json:"transaction_id"`
	Qty           int `json:"qty"`
	SubAmount     int `json:"sub_amount"`
}
