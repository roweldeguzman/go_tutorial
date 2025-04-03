package service

import (
	"api/models"
	"api/repository"
)

type PostsService struct {
	repo *repository.PostsRepository
}

func NewPostsService(repo *repository.PostsRepository) *PostsService {
	return &PostsService{repo}
}

func (s *PostsService) Create(post *models.Posts) (*models.Posts, error) {
	return s.repo.Create(post)
}
