package repositories

import (
	"waysbeans_be/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindProfile(UserID int) ([]models.Profile, error)
	GetProfile(ID int) (models.Profile, error)
	CreateProfile(profile models.Profile) (models.Profile, error)
	UpdateProfile(profile models.Profile) (models.Profile, error)
	DeleteProfile(profile models.Profile) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProfile(UserID int) ([]models.Profile, error) {
	var profiles []models.Profile
	err := r.db.Preload("User").Find(&profiles, "user_id = ?", UserID).Error

	return profiles, err
}

func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").First(&profile, ID).Error

	return profile, err
}

func (r *repository) CreateProfile(profile models.Profile) (models.Profile, error) {
	err := r.db.Preload("User").Create(&profile).Error

	return profile, err
}

func (r *repository) UpdateProfile(profile models.Profile) (models.Profile, error) {
	err := r.db.Preload("User").Save(&profile).Error

	return profile, err
}

func (r *repository) DeleteProfile(profile models.Profile) (models.Profile, error) {
	err := r.db.Preload("User").Delete(&profile).Error

	return profile, err
}
