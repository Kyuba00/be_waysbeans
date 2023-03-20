package repositories

import (
	"waysbeans_be/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCarts() ([]models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart) (models.Cart, error)
	GetTransactionID(ID int) (models.Transaction, error)
	GetIDTransaction(UserID int) (models.Transaction, error)
	FindCartsByTransID(TrsID int) ([]models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCarts() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Find(&carts).Error

	return carts, err
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Product").First(&cart, ID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Preload("Product").Create(&cart).Error

	return cart, err
}

func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Preload("Product").Save(&cart).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Preload("Product").Delete(&cart).Error

	return cart, err
}

func (r *repository) GetTransactionID(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Find(&transaction, "user_id = ? AND status = ?", ID, "Active").Error

	return transaction, err
}

func (r *repository) GetIDTransaction(UserID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Find(&transaction, "user_id =? AND status = ?", UserID, "Active").Error

	return transaction, err
}

func (r *repository) FindCartsByTransID(TrsID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Find(&carts, "transaction_id = ?", TrsID).Error

	return carts, err
}
