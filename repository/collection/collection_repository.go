package repository

import (
	"ta13-svc/entity"
)

type CollectionRepository interface {
	Insert(collection entity.Collection)
	FindAll() (collections []entity.Collection)
}
