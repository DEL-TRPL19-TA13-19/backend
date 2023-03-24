package repository

import (
	"ta13-svc/entity"
)

type UserRepository interface {
	GetUser(user entity.User)
	GetAllUsers(users []entity.User)
}
