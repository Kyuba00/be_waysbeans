package repositories

import (
	"waysbeans_be/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUSer(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Transaction").Preload("Transaction.Cart").Find(&users).Error

	return users, err
}

func (r *repository) GetUSer(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Transaction").Preload("Transaction.Cart").First(&user, ID).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Preload("Transaction").Preload("Transaction.Cart").Create(&user).Error

	return user, err
}

func (r *repository) UpdateUSer(user models.User) (models.User, error) {
	err := r.db.Preload("Transaction").Preload("Transaction.Cart").Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Preload("Transaction").Preload("Transaction.Cart").Delete(&user).Error

	return user, err
}
