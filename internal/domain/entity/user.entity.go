package entity

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	FirstName string    `gorm:"type:varchar(100);not null"`
	LastName  string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	Token     string    `gorm:"type:varchar(100);not null"`
	CreatedAt string    `gorm:"type:timestamp;not null"`
	UpdatedAt string    `gorm:"type:timestamp;not null"`
}

func (u *User) TableName() string {
	return "users"
}
