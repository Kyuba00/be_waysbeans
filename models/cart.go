package models

type Cart struct {
	ID            int         `json:"id" gorm:"primary_key:auto_increment"`
	ProductID     int         `json:"product_id" gorm:"type: int"`
	Product       Product     `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TransactionID int         `json:"transaction_id" gorm:"type: int"`
	Transaction   Transaction `json:"-"`
	Qty           int         `json:"qty" gorm:"type: int"`
	SubAmount     int         `json:"sub_amount" gorm:"type: int"`
}
