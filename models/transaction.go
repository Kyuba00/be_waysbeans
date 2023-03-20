package models

type Transaction struct {
	ID     int                 `json:"id" gorm:"primary_key:auto_increment"`
	Status string              `json:"status" gorm:"type: varchar(255)"`
	UserID int                 `json:"user_id"`
	User   UserProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cart   []Cart              `json:"cart" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount int                 `json:"amount" gorm:"type: int"`
}
