package repository

import "github.com/isaquecsilva/posts-rest-api/core/entities"

type Repository interface {
	Insert()
	Find(id int) entities.Post
	FindAll() []entities.Post
}
