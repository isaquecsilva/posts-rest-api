package usecase

import (
	"github.com/isaquecsilva/posts-rest-api/core/entities"
	"github.com/isaquecsilva/posts-rest-api/core/repository"
)

type PostServiceUseCase interface {
	CreatePost(post entities.Post)
	GetPost(id int) (entities.Post, error)
	GetAllPosts() ([]entities.Post, error)
}

type PostService struct {
	repo    repository.Repository
	maxSize int
}

func NewPostService(repo repository.Repository, maxSize int) *PostService {
	return &PostService{
		repo,
		maxSize,
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
