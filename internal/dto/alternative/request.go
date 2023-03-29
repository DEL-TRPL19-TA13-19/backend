package dto

import (
	"ta13-svc/internal/entity"
)

type AlternativeGetByCollectionIDRequest struct {
	CollectionID string `json:"collection_id" param:"collection_id" validate:"required"`
}

type AlternativeGetByIDRequest struct {
	ID string `json:"id" param:"id" validate:"required"`
}

type AlternativeCreateRequest struct {
	entity.AlternativeEntity
	CollectionID string `json:"collection_id"`
}

type AlternativeUpdateRequest struct {
	ID string `param:"id" validate:"required"`
	entity.AlternativeEntity
}

type AlternativeDeleteRequest struct {
	ID string `param:"id" validate:"required"`
	entity.AlternativeEntity
}
