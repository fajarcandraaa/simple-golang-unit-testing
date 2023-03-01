package app

import (
	"github.com/fajarcandraaa/simple-golang-unit-testing/entity/userentity"
)

// SetMigrationTable is used to register entity model which want to be migrate
func SetMigrationTable() []interface{} {
	var migrationData = []interface{}{
		&userentity.User{},
	}

	return migrationData
}
