package repository

import (
	"ta13-svc/internal/entity"
)

type CollectionRepository interface {
	Insert(collection entity.Collection)
	FindAll() (collections []entity.Collection)
}
