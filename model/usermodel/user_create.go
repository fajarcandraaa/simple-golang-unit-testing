package usermodel

import (
	"context"

	"github.com/pkg/errors"

	"github.com/fajarcandraaa/simple-golang-unit-testing/config/app"
	"github.com/fajarcandraaa/simple-golang-unit-testing/entity"
	"github.com/fajarcandraaa/simple-golang-unit-testing/entity/userentity"
)

// SaveNewUser is used to run query insert
func (r *UserRepository) SaveNewUser(ctx context.Context, payload *userentity.User) error {

	err := r.db.Create(payload).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return entity.ErrUserNotExist
		case app.ErrUniqueViolation:
			return entity.ErrUserAlreadyExist
		default:
			return errors.Wrap(parsed, "build statement query to insert user from database")
		}
	}
	return nil
}
