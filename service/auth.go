package service

import (
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/io_models"
)

// GetUserByEmail ...
func GetUserByEmail(email string) (*io_models.Register, error) {
	db := database.GetDB()

	var auth io_models.Register

	if err := db.Where("email = ? ", email).First(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}
