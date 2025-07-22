package entity

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// UserID represents a unique identifier for a user
// This abstraction allows different storage implementations
type UserID string

// User represents our core domain entity, completely independent of any framework
// This entity focuses purely on business logic and domain rules
type User struct {
	ID              UserID
	Name            string
	Designation     *string
	Email           string
	EmailVerifiedAt *time.Time
	Password        string
	RememberToken   *string
	// Additional fields for extended profile
	FirstName string
	LastName  string
	Phone     *string
	Address   *string
	// Audit fields
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// NewUser creates a new user with validation
func NewUser(name, email, password string) (*User, error) {
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	
	if err := validateName(name); err != nil {
		return nil, err
	}
	
	if err := validatePassword(password); err != nil {
		return nil, err
	}
	
	now := time.Now()
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// UpdateProfile updates user profile information
func (u *User) UpdateProfile(firstName, lastName, phone, address string) {
	u.FirstName = firstName
	u.LastName = lastName
	u.Phone = &phone
	u.Address = &address
	u.Name = strings.TrimSpace(firstName + " " + lastName)
	u.UpdatedAt = time.Now()
}

// VerifyEmail marks the user's email as verified
func (u *User) VerifyEmail() {
	now := time.Now()
	u.EmailVerifiedAt = &now
	u.UpdatedAt = now
}

// ChangePassword changes user password with validation
func (u *User) ChangePassword(newPassword string) error {
	if err := validatePassword(newPassword); err != nil {
		return err
	}
	u.Password = newPassword
	u.UpdatedAt = time.Now()
	return nil
}

// GetFullName returns the full name of the user
func (u *User) GetFullName() string {
	return strings.TrimSpace(u.Name)
}

// IsValid checks if the user entity is valid
func (u *User) IsValid() bool {
	return u.Name != "" && u.Email != "" && u.Password != ""
}

// Domain validation functions
func validateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	if !strings.Contains(email, "@") {
		return errors.New("invalid email format")
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	if len(name) < 2 {
		return errors.New("name must be at least 2 characters")
	}
	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return errors.New("password is required")
	}
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}

// ParseUserIDToUint converts UserID to uint for database operations
func ParseUserIDToUint(id UserID) (uint, error) {
	if id == "" {
		return 0, errors.New("empty user ID")
	}
	
	parsedID, err := strconv.ParseUint(string(id), 10, 32)
	if err != nil {
		return 0, errors.New("invalid user ID format")
	}
	
	return uint(parsedID), nil
}
