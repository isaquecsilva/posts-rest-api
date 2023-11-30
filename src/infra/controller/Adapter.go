package controller

import "github.com/isaquecsilva/posts-rest-api/core/entities"

type Adapter interface {
	GetIdParam() (int, error)
	GetRequest() (entities.Post, error)
	SendResponse(code int, obj any) error
	Validate() error
}
