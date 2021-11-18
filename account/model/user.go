package model

import (
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
}

type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
}

type User struct {
	UID      uuid.UUID `db:"uid" json:"uid"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"-"`
	Name     string    `db:"name" json:"name"`
	ImageURL string    `db:"name" json:"imageUrl"`
	Website  string    `db:"website" json:"website"`
}
