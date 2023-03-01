package test

import (
	"context"
	"log"
	"testing"

	"github.com/fajarcandraaa/simple-golang-unit-testing/test/faker"
	"github.com/fajarcandraaa/simple-golang-unit-testing/test/seeders"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByID(t *testing.T) {
	ctx := context.Background()
	err := refreshTable()
	if err != nil {
		log.Fatalf("Error user refreshing table %v\n", err)
	}
	s := userRepository()
	_, err = seeders.SeedUser(server.DB)
	if err != nil {
		log.Fatalf("Error user refreshing table %v\n", err)
	}
	t.Run("if document body request fulfilled, it should not return an error", func(t *testing.T) {
		fakeFindUser, err := s.FindUser(ctx, faker.UserID001)
		assert.NoError(t, err)
		assert.NotNil(t, fakeFindUser)
	})

	t.Run("if given invalid type UUID, it should return an error", func(t *testing.T) {
		fakeUserID := "this is not uuid"
		fakeFindUser, err := s.FindUser(ctx, fakeUserID)
		assert.Error(t, err)
		assert.Nil(t, fakeFindUser)
	})
}
