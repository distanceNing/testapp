package repo

import "time"

type UserLogin struct {
	UserId       string `gorm:"primarykey"`
	UserType     int
	UserPassword string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
