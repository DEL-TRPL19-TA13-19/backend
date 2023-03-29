package dto

import (
	"ta13-svc/internal/entity"
	"ta13-svc/pkg/response"
)

type CollectionsGetResponse struct {
	Datas []entity.CollectionEntityModel
}
type CollectionsGetResponseDoc struct {
	Body struct {
		Meta response.Meta          `json:"meta"`
		Data CollectionsGetResponse `json:"data"`
	} `json:"body"`
}

type CollectionGetByIDResponse struct {
	Datas entity.CollectionEntityModel
}
type CollectionGetByIDResponseDoc struct {
	Body struct {
		Meta response.Meta             `json:"meta"`
		Data CollectionGetByIDResponse `json:"data"`
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
	ID *string `json:"id"`
}
type CollectionDeleteResponseDoc struct {
	Body struct {
		Meta response.Meta            `json:"meta"`
		Data CollectionDeleteResponse `json:"data"`
	} `json:"body"`
}
