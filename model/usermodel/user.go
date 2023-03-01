package usermodel

import (
	"github.com/fajarcandraaa/simple-golang-unit-testing/helpers/errorcodehandling"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db        *gorm.DB
	codeError *errorcodehandling.CodeError
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
