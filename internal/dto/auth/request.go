package dto

import "ta13-svc/internal/entity"

type AuthLoginRequest struct {
	Username string `json:"username" validate:"required" example:"nathan"`
	Password string `json:"password" validate:"required" example:"pass1234"`
}

type AuthRegisterRequest struct {
	entity.UserEntity
}
