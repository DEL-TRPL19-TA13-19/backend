package dto

import "ta13-svc/internal/entity"

type AuthLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthRegisterRequest struct {
	entity.UserEntityModel
}
