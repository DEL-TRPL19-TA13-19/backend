package dto

import (
	"ta13-svc/internal/entity"
)

type AlternativeGetByCollectionIDRequest struct {
	CollectionID string `json:"id"`
}

type AlternativeCreateRequest struct {
	entity.AlternativeEntity
	CollectionID string `json:"collection_id"`
}

type AlternativeUpdateRequest struct {
	ID string `param:"id" validate:"required"`
	entity.AlternativeEntityModel
}

type AlternativeDeleteRequest struct {
	ID string `param:"id" validate:"required"`
	entity.AlternativeEntityModel
}
