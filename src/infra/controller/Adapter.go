package controller

import "github.com/isaquecsilva/posts-rest-api/core/entities"

type Adapter interface {
	GetID() (int, error)
	GetRequest() (entities.Post, error)
	SendResponse(code int, obj any) error
	Validate() error
}
