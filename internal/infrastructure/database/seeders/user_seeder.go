package seeders

import (
	"time"

	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"
	"kossti/pkg/hash"

	"gorm.io/gorm"
)

// UserSeeder handles seeding admin/user accounts
type UserSeeder struct {
	BaseSeeder
}

// NewUserSeeder creates a new user seeder
func NewUserSeeder() *UserSeeder {
	return &UserSeeder{BaseSeeder: BaseSeeder{name: "Users"}}
}

// Seed implements the Seeder interface
func (us *UserSeeder) Seed(db *gorm.DB) error {
	// Admin credentials (from request)
	name := "admin"
	email := "admin@gmail.com"
	rawPassword := "admin@*123"

	// Hash the password
	hashed, err := hash.HashPassword(rawPassword)
	if err != nil {
		return err
	}

	userEntity := &entities.User{
		Name:            name,
		Email:           email,
		Password:        hashed,
		EmailVerifiedAt: func() *time.Time { t := time.Now(); return &t }(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	var userModel models.UserModel

	if err := db.Where("email = ?", email).First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			userModel.FromEntity(userEntity)
			if err := db.Create(&userModel).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
