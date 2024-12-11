package user_usecase

import (
	"context"

	"github.com/dpcamargo/fullcycle-auction/internal/entity/user_entity"
	"github.com/dpcamargo/fullcycle-auction/internal/internal_error"
)

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface 	
}

type UserOutputDTO struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserUseCaseInterface interface {
}

func (u *UserUseCase) FindUserByID(ctx context.Context,	id string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := u.UserRepository.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &UserOutputDTO{
		ID:   userEntity.ID,
		Name: userEntity.Name,
	}, nil
}
