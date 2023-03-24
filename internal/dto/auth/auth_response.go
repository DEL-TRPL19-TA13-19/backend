package dto

import (
	"ta13-svc/internal/entity"
	"ta13-svc/pkg/response"
)

type AuthLoginResponse struct {
	Message string `json:"message"`
	entity.UserEntityModel
}
type AuthLoginResponseDoc struct {
	Body struct {
		Meta response.Meta `json:"meta"`
		Data AuthLoginResponse
	} `json:"body"`
}

type GetUsersResponse struct {
	Users []entity.UserEntityModel `json:"users"`
}
type GetUsersResponseDoc struct {
	Body struct {
		Meta response.Meta `json:"meta"`
		Data GetUsersResponse
	} `json:"body"`
}

type AuthRegisterResponse struct {
	entity.UserEntityModel
}

type AuthRegisterResponseDoc struct {
	Body struct {
		Meta response.Meta        `json:"meta"`
		Data AuthRegisterResponse `json:"data"`
	} `json:"body"`
}
