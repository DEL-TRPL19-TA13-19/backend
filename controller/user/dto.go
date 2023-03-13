package user

import "ta13-svc/model"

type ResponseGetUsers struct {
	Users []model.User `json:"users"`
}

type ResponseHello struct {
	Message string `json:"message"`
}
