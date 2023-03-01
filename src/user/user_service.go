package user

import (
	"context"

	"github.com/fajarcandraaa/simple-golang-unit-testing/entity/userentity"
	repositories "github.com/fajarcandraaa/simple-golang-unit-testing/model"
)

type service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) FindUser(ctx context.Context, userID string) (*userentity.FindUser, error) {

	user, err := s.repo.User.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	resultUser := userentity.FindUser{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}

	return &resultUser, nil
}
