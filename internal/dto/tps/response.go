package dto

import (
	"ta13-svc/internal/entity"
	"ta13-svc/pkg/response"
)

type TpsGetResponse struct {
	Datas []entity.TpsEntityModel
}
type TpsGetResponseDoc struct {
	Body struct {
		Meta response.Meta           `json:"meta"`
		Data []entity.TpsEntityModel `json:"data"`
	} `json:"body"`
}

type TpsGetByIdResponse struct {
	entity.TpsEntityModel
}
type TpsGetByIdResponseDoc struct {
	Body struct {
		Meta response.Meta      `json:"meta"`
		Data TpsGetByIdResponse `json:"data"`
	} `json:"body"`
}

type TpsCreateResponse struct {
	entity.TpsEntityModel
}
type TpsCreateResponseDoc struct {
	Body struct {
		Meta response.Meta     `json:"meta"`
		Data TpsCreateResponse `json:"data"`
	} `json:"body"`
}

type TpsUpdateResponse struct {
	entity.TpsEntityModel
}
type TpsUpdateResponseDoc struct {
	Body struct {
		Meta response.Meta     `json:"meta"`
		Data TpsUpdateResponse `json:"data"`
	} `json:"body"`
}

type TpsDeleteResponse struct {
	ID string `json:"id"`
}
type TpsDeleteResponseDoc struct {
	Body struct {
		Meta response.Meta     `json:"meta"`
		Data TpsDeleteResponse `json:"data"`
	} `json:"body"`
}
