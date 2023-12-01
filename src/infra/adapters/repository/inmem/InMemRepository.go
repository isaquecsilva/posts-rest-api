package inmem

import (
	"errors"

	"github.com/isaquecsilva/posts-rest-api/core/entities"
)

type DatabaseMock struct {
	Posts []entities.Post
}

type RepositoryMock struct {
	DatabaseMock
}

func (rm *RepositoryMock) Insert(post entities.Post) error {
	rm.DatabaseMock.Posts = append(rm.DatabaseMock.Posts, post)
	return nil
}

func (rm *RepositoryMock) Find(id int) (entities.Post, error) {
	if len(rm.DatabaseMock.Posts) < id || id == 0 {
		return entities.Post{}, errors.New("post not found")
	}

	return rm.DatabaseMock.Posts[id-1], nil
}

func (rm *RepositoryMock) FindAll() ([]entities.Post, error) {
	return rm.DatabaseMock.Posts, nil
}
