package usecase

import (
	"github.com/isaquecsilva/posts-rest-api/core/entities"
	"github.com/isaquecsilva/posts-rest-api/core/repository"
)

type PostServiceUseCase interface {
	CreatePost(post entities.Post) error
	GetPost(id int) (entities.Post, error)
	GetAllPosts() ([]entities.Post, error)
}

type PostService struct {
	repo repository.Repository
}

func NewPostService(repo repository.Repository) *PostService {
	return &PostService{
		repo,
	}
}

func (ps *PostService) CreatePost(post entities.Post) error {
	return ps.repo.Insert(post)
}

func (ps *PostService) GetPost(id int) (entities.Post, error) {
	return ps.repo.Find(id)
}

func (ps *PostService) GetAllPosts() ([]entities.Post, error) {
	return ps.repo.FindAll()
}
