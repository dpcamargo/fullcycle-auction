package user_entity

import (
	"context"

	"github.com/dpcamargo/fullcycle-auction/internal/internal_error"
)

type User struct {
	ID   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserByID(
		ctx context.Context, id string) (*User, *internal_error.InternalError)
}
