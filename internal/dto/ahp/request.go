package dto

import "ta13-svc/internal/entity"

type AHPByCollectionIDRequest struct {
	CollectionID string `json:"collection_id" param:"collection_id" validate:"required"`
}

type CriteriaAlternativeUpdateRequest struct {
	Pairwise entity.Matrix `json:"pairwise"`
}
