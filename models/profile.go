package models

import "time"

type Profile struct {
	ID        int                 `json:"id" gorm:"primary_key:auto_increment"`
	Image     string              `json:"image" gorm:"type: varchar(255)"`
	Address   string              `json:"address" gorm:"type: varchar(255)"`
	Postcode  string              `json:"postcode" gorm:"type: varchar(255)"`
	Phone     string              `json:"phone" gorm:"type: varchar(255)"`
	UserID    int                 `json:"user_id"`
	User      UserProfileResponse `json:"user"`
	CreatedAt time.Time           `json:"-"`
	UpdatedAt time.Time           `json:"-"`
}

type ProfileUserResponse struct {
	Image    string `json:"image"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	Phone    string `json:"phone"`
}

type ProfileTransactionResponse struct {
	UserID   int    `json:"user_id"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	Phone    string `json:"phone"`
}

func (ProfileUserResponse) TableName() string {
	return "profiles"
}

func (ProfileTransactionResponse) TableName() string {
	return "profiles"
}
