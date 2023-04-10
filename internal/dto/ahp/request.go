package dto

type AHPByCollectionIDRequest struct {
	CollectionID string `json:"collection_id" param:"collection_id" validate:"required"`
}
