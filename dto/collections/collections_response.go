package collections

import "ta13-svc/entity"

type ResponseGetCollections struct {
	Collections []entity.Collection `json:"collections"`
}
