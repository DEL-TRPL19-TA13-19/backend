package dto

import (
	"ta13-svc/internal/entity"
)

type CollectionGetByIDRequest struct {
	ID string `json:"id" param:"id" validate:"required"`
}

type CollectionsGetByUserIDRequest struct {
	UserID string `param:"user_id" validate:"required"`
}

type CollectionCreateRequest struct {
	entity.CollectionEntity
}

type CollectionUpdateRequest struct {
	ID string `param:"id" validate:"required"`
	entity.CollectionEntity
}

type CollectionDeleteRequest struct {
	ID string `param:"id" validate:"required"`
	entity.CollectionEntity
}
