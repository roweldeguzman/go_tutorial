package posts

import "api/service"

type PostsController struct {
	service *service.PostsService
}

func NewPostController(service *service.PostsService) *PostsController {

	return &PostsController{service}
}
