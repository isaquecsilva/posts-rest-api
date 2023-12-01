package controller

import (
	"fmt"
	"net/http"

	"github.com/isaquecsilva/posts-rest-api/config/rest_errors"
	"github.com/isaquecsilva/posts-rest-api/core/usecase"
)

type Controller struct {
	postService usecase.PostServiceUseCase
}

func NewController(postService usecase.PostServiceUseCase) *Controller {
	return &Controller{
		postService,
	}
}

func (c *Controller) CreateNewPost(adapter Adapter) {
	// Whether users request fields are invalid
	// Then, we return a bad request rest error message.
	if err := adapter.Validate(); err != nil {
		badRequestErrorMessage := rest_errors.NewBadRequestRestError(err.Error())
		adapter.SendResponse(badRequestErrorMessage.Code, badRequestErrorMessage)
		return
	}

	post, err := adapter.GetRequest()

	// We couldn't parse the request?
	// So, we return an internal server rest error message.
	if err != nil {
		fmt.Println("chegou aqui")
		restErrorMessage := rest_errors.NewInternalServerRestError(err.Error())
		adapter.SendResponse(restErrorMessage.Code, restErrorMessage)
		return
	}

	// If everything is ok, we call the post service
	if err = c.postService.CreatePost(post); err == nil {
		adapter.SendResponse(http.StatusOK, "Post created.")
	} else {
		internalError := rest_errors.NewInternalServerRestError(err.Error())
		adapter.SendResponse(internalError.Code, internalError)
	}

}

func (c *Controller) GetPost(adapter Adapter) {
	// Whether the postId param couldn't be
	// get, we receive the error message and send it
	// as a response to final user.
	postId, err := adapter.GetID()

	if err != nil {
		badRequest := rest_errors.NewBadRequestRestError(err.Error())
		adapter.SendResponse(badRequest.Code, badRequest)
		return
	}

	if post, err := c.postService.GetPost(postId); err != nil {
		internalError := rest_errors.NewInternalServerRestError(err.Error())
		adapter.SendResponse(internalError.Code, internalError)
	} else {
		adapter.SendResponse(http.StatusOK, post)
	}
}

func (c *Controller) GetAllPosts(adapter Adapter) {
	if posts, err := c.postService.GetAllPosts(); err == nil {
		adapter.SendResponse(http.StatusOK, posts)
	} else {
		internalError := rest_errors.NewInternalServerRestError(err.Error())
		adapter.SendResponse(internalError.Code, internalError)
	}
}
