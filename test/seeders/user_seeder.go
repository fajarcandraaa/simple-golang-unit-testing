package seeders

import (
	"github.com/fajarcandraaa/simple-golang-unit-testing/entity/userentity"
	"github.com/fajarcandraaa/simple-golang-unit-testing/test/faker"
	"github.com/jinzhu/gorm"
)

// SeedUser is seeder for testing database
func SeedUser(db *gorm.DB) (*userentity.User, error) {
	fakeuser := faker.FakeUser()
	err := db.Create(&fakeuser).Error
	if err != nil {
		return nil, err
	}

	return fakeuser, nil
}
