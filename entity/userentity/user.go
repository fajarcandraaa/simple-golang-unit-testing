package userentity

import (
	"time"
)

// User -> initialization user entity
type User struct {
	ID        string    `gorm:"size:36;not null;unique index;primaryKey"`
	Firstname string    `gorm:"size:255;"`
	Lastname  string    `gorm:"size:255;"`
	Phone     string    `gorm:"size:50;"`
	Avatar    string    `gorm:"size:255;"`
	Email     string    `gorm:"size:100;unique"`
	Username  string    `gorm:"size:100;unique"`
	Password  string    `gorm:"size:100;"`
	Status    string    `gorm:"size:50;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// FindUser is struct to handle respone while find user by ID
type FindUser struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Phone     string `json:"phone_number"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
