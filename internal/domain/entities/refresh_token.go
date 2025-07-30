package entities

type RefreshToken struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index;not null"`
	Token     string `gorm:"not null;unique"`
	ExpiresAt int64  `gorm:"not null"`
}
