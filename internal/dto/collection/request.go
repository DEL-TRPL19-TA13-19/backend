package dto

import (
	"ta13-svc/internal/entity"
)

type CollectionGetByUserIDRequest struct {
	UserID string `json:"id"`
}

type CollectionCreateRequest struct {
	entity.CollectionEntity
}

type CollectionUpdateRequest struct {
	ID string `param:"id" validate:"required"`
	entity.CollectionEntityModel
}

type CollectionDeleteRequest struct {
	ID string `param:"id" validate:"required"`
	entity.CollectionEntityModel
}
