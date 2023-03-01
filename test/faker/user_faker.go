package faker

import (
	"time"

	"github.com/fajarcandraaa/simple-golang-unit-testing/entity/userentity"
)

const (
	UserID001 = "a61a88bd-3c0e-4c5b-a203-40e92a22b3ca"
	UserFirstName001 = "So"
	UserLastName001 = "Klin"
	UserPhoneNumber001 = "08123456789"
	UserAvatar001 = "avatar01"
	UserEmail001 = "user001@gmail.com"
	UserUsername001 = "user01"
	UserPassword001 = "123456"
	UserStatus001 = "active"
	
	UserID002 = "463234fc-b2af-11ed-afa1-0242ac120002"
	UserFirstName002 = "Rin"
	UserLastName002 = "So"
	UserPhoneNumber002 = "08123456788"
	UserAvatar002 = "avatar02"
	UserEmail002 = "user002@gmail.com"
	UserUsername002 = "user02"
	UserPassword002 = "123456"
	UserStatus002 = "active"

	UserID003 = "4f45a0ec-b2af-11ed-afa1-0242ac120002"
	UserFirstName003 = "Good"
	UserLastName003 = "Day"
	UserPhoneNumber003 = "08123456789"
	UserAvatar003 = "avatar03"
	UserEmail003 = "user003@gmail.com"
	UserUsername003 = "user03"
	UserPassword003 = "123456"
	UserStatus003 = "active"
)

// FakeUser initiate fake data user
func FakeUser() *userentity.User {
	t := time.Now()
	fakeUser := &userentity.User{
		ID:        UserID001,
		Firstname: UserFirstName001,
		Lastname:  UserLastName001,
		Phone:     UserPhoneNumber001,
		Avatar:    UserAvatar001,
		Email:     UserEmail001,
		Username:  UserUsername001,
		Password:  UserID001,
		Status:    UserID001,
		CreatedAt: t,
		UpdatedAt: t,
	}

	return fakeUser
}
