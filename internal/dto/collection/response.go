package dto

import (
	"ta13-svc/internal/entity"
	"ta13-svc/pkg/response"
)

type CollectionGetByUserIDResponse struct {
	Datas []entity.CollectionEntityModel
}
type CollectionGetByUserIDResponseDoc struct {
	Body struct {
		Meta response.Meta                 `json:"meta"`
		Data CollectionGetByUserIDResponse `json:"data"`
	} `json:"body"`
}

type CollectionCreateResponse struct {
	entity.CollectionEntityModel
}
type CollectionCreateResponseDoc struct {
	Body struct {
		Meta response.Meta            `json:"meta"`
		Data CollectionCreateResponse `json:"data"`
	} `json:"body"`
}

type CollectionUpdateResponse struct {
	entity.CollectionEntityModel
}
type CollectionUpdateResponseDoc struct {
	Body struct {
		Meta response.Meta            `json:"meta"`
		Data CollectionUpdateResponse `json:"data"`
	} `json:"body"`
}

type CollectionDeleteResponse struct {
	entity.CollectionEntityModel
}
type CollectionDeleteResponseDoc struct {
	Body struct {
		Meta response.Meta            `json:"meta"`
		Data CollectionDeleteResponse `json:"data"`
	} `json:"body"`
}
