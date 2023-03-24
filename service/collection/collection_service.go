package collection

import "ta13-svc/dto/collections"

type CollectionService interface {
	Create()
	List(response []collections.ResponseGetCollections)
}
