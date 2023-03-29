package dto

import (
	"ta13-svc/internal/entity"
	"ta13-svc/pkg/response"
)

type AlternativesGetResponse struct {
	Datas []entity.AlternativeEntityModel
}
type AlternativesGetResponseDoc struct {
	Body struct {
		Meta response.Meta           `json:"meta"`
		Data AlternativesGetResponse `json:"data"`
	} `json:"body"`
}

type AlternativeGetByIDResponse struct {
	Datas entity.AlternativeEntityModel
}
type AlternativeGetByIDResponseDoc struct {
	Body struct {
		Meta response.Meta              `json:"meta"`
		Data AlternativeGetByIDResponse `json:"data"`
	} `json:"body"`
}

type AlternativeGetByCollectionIDResponse struct {
	Datas []entity.AlternativeEntityModel
}
type AlternativeGetByCollectionIDResponseDoc struct {
	Body struct {
		Meta response.Meta                        `json:"meta"`
		Data AlternativeGetByCollectionIDResponse `json:"data"`
	} `json:"body"`
}

type AlternativeCreateResponse struct {
	entity.AlternativeEntityModel
}
type AlternativeCreateResponseDoc struct {
	Body struct {
		Meta response.Meta             `json:"meta"`
		Data AlternativeCreateResponse `json:"data"`
	} `json:"body"`
}

type AlternativeUpdateResponse struct {
	entity.AlternativeEntityModel
}
type AlternativeUpdateResponseDoc struct {
	Body struct {
		Meta response.Meta             `json:"meta"`
		Data AlternativeUpdateResponse `json:"data"`
	} `json:"body"`
}

type AlternativeDeleteResponse struct {
	ID *string `json:"id"`
}
type AlternativeDeleteResponseDoc struct {
	Body struct {
		Meta response.Meta             `json:"meta"`
		Data AlternativeDeleteResponse `json:"data"`
	} `json:"body"`
}
