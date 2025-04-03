package repository

import "api/models"

type PostsRepository struct {
}

func NewPostsRepository() *UsersRepository {
	return &UsersRepository{}
}

func (r *PostsRepository) Create(post *models.Posts) (*models.Posts, error) {
	return nil, nil
}
