package model

import (
	"context"

	"github.com/fajarcandraaa/simple-golang-unit-testing/entity/userentity"
	"github.com/fajarcandraaa/simple-golang-unit-testing/model/usermodel"
	"github.com/jinzhu/gorm"
)

type User interface {
	FindUserByID(ctx context.Context, userID string) (*userentity.User, error)
}

func NewUser(db *gorm.DB) User {
	return usermodel.NewUserRepository(db)
}
