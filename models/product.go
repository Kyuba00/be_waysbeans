package models

import "time"

type Product struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Name        string    `json:"name" gorm:"type: varchar(255)"`
	Stock       int       `json:"stock" gorm:"type: int"`
	Price       int       `json:"price" gorm:"type: int"`
	Description string    `json:"description" gorm:"type: text"`
	Image       string    `json:"image" gorm:"type: varchar(255)"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
