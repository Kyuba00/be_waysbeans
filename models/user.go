package models

import "time"

type User struct {
	ID            int         `json:"id" gorm:"primary_key:auto_increment"`
	Name          string      `json:"name" gorm:"type: varchar(255)"`
	Email         string      `json:"email" gorm:"type: varchar(255)"`
	Password      string      `json:"password" gorm:"type: varchar(255)"`
	Status        string      `json:"status" gorm:"type: varchar(255)"`
	TransactionID int         `json:"transaction_id"`
	Transaction   Transaction `json:"transaction"`
	CreatedAt     time.Time   `json:"-"`
	UpdatedAt     time.Time   `json:"-"`
}

type UserProfileResponse struct {
	ID    int    `json:"ID" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func (UserProfileResponse) TableName() string {
	return "users"
}
