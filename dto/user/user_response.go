package user

import "ta13-svc/entity"

type ResponseGetUsers struct {
	Users []entity.User `json:"users"`
}

type ResponseHello struct {
	Message string `json:"message"`
}
