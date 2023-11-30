package repository

import "github.com/isaquecsilva/posts-rest-api/core/entities"

type Repository interface {
	Insert(post entities.Post) error
	Find(id int) (entities.Post, error)
	FindAll() ([]entities.Post, error)
}
