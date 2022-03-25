package entity

import "github.com/google/uuid"

type SiginUp struct {
	Name     string `json:"name" binging:"required"`
	Email    string `json:"email" binging:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type RegisterTask struct {
	OwnerId     uuid.UUID `json:"user_id" binding:"required"`
	Title       string    `json:"title" binging:"required"`
	Description string    `json:"description" binging:"required"`
}

type JSONMessage struct {
	Message string `json:"message"`
}
