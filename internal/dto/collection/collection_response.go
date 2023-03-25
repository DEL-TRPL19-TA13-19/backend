package collection

import (
	"ta13-svc/internal/entity"
)

type ResponseGetCollections struct {
	Collections []entity.Collection `json:"collection"`
}
