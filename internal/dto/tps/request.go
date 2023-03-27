package dto

import (
	"ta13-svc/internal/entity"
)

type TpsGetRequest struct {
	entity.TpsEntityModel
}

type TpsGetByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type TpsCreateRequest struct {
	entity.TpsEntity
}

type TpsUpdateRequest struct {
	ID string `param:"id" validate:"required"`
	entity.TpsEntity
}

type TpsDeleteRequest struct {
	ID string `param:"id" validate:"required"`
	entity.TpsEntity
}
