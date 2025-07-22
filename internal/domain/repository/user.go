package repository

// UserRepository defines the contract for user data access.
// - This interface is part of the Domain Layer in Clean Code Architecture.
// - It uses the Repository Pattern to abstract data storage details from business logic.
// - Enables Dependency Injection: use cases depend on this interface, not a concrete DB.
// - Makes the code testable and decoupled from infrastructure (e.g., Postgres, MongoDB, etc).

import (
	"context"
	"kossti/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	GetByID(ctx context.Context, id uint) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*entities.User, error)
	
	// Additional methods for better user management
	Count(ctx context.Context) (int64, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*entities.User, error)
	GetByEmailOrUsername(ctx context.Context, emailOrUsername string) (*entities.User, error)
}
