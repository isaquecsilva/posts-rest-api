package usecase

import (
	"github.com/isaquecsilva/posts-rest-api/core/entities"
	"github.com/isaquecsilva/posts-rest-api/core/repository"
)

type PostServiceUseCase interface {
	CreatePost(post entities.Post)
	GetPost(id int) entities.Post
	GetAllPosts() []entities.Post
}

type PostService struct {
	repo repository.Repository
}

func NewPostService(repo repository.Repository) *PostService {
	return &PostService{
		repo,
	}
}

func (ps *PostService) CreatePost(post entities.Post) {

}

func (ps *PostService) GetPost(id int) entities.Post {
	return entities.NewPost("", "")
}

func (ps *PostService) GetAllPosts() []entities.Post {
	return []entities.Post{}
}
