package user

import (
	"context"

	"github.com/fajarcandraaa/simple-golang-unit-testing/entity/userentity"
)

type Service interface {
	FindUser(ctx context.Context, userID string) (*userentity.FindUser, error)
}
