package entity

import (
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-todolist/pkg/auth"
)

type User struct {
	Id           uuid.UUID `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(200)" json:"name"`
	Email        string    `gorm:"type:varchar(200)" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255)" json:"password_hash"`
}

func NewUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, ErrMissingName
	}

	if email == "" {
		return nil, ErrMissingEmail
	}

	if password == "" {
		return nil, ErrMissingPassword
	}

	hash, err := auth.GeneratePasswordHash(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: hash,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	return auth.CheckPasswordHash(password, u.PasswordHash)
}
