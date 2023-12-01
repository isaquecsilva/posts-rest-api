package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/isaquecsilva/posts-rest-api/config"
	"github.com/isaquecsilva/posts-rest-api/core/entities"

	validator "github.com/go-playground/validator/v10"
)

type HTTPPost struct {
	NickName string `validate:"required,min=5,max=20"`
	Content  string `validate:"required,min=1,max=255"`
}

type HttpAdapter struct {
	request  *http.Request
	response http.ResponseWriter
	post     entities.Post
}

func NewHttpAdapter(request *http.Request, response http.ResponseWriter) *HttpAdapter {
	return &HttpAdapter{request: request, response: response}
}

func (ha *HttpAdapter) GetID() (int, error) {
	id, err := strconv.ParseInt(
		ha.request.URL.Query().Get("id"),
		10,
		32,
	)

	return int(id), err
}

func (ha *HttpAdapter) GetRequest() (entities.Post, error) {
	if ha.post.NickName == "" || ha.post.Content == "" {
		return ha.post, errors.New("Post object not parsed")
	}

	return ha.post, nil
}

func (ha *HttpAdapter) SendResponse(code int, obj any) error {

	buf, err := json.Marshal(obj)

	if err != nil {
		return err
	}

	ha.response.WriteHeader(code)
	ha.response.Write(buf)

	return nil
}

func (ha *HttpAdapter) Validate() error {
	length := ha.request.ContentLength

	// Checking request body length
	if length != -1 && length > config.MAX_ALLOWED_REQUEST_BODY_LENGTH {
		return fmt.Errorf("Body size overtake the limit allowed.")
	}
	defer ha.request.Body.Close()
	buf, _ := io.ReadAll(ha.request.Body)

	// Parsing request body
	var post HTTPPost
	var err error = json.Unmarshal(buf, &post)

	if err != nil {
		return err
	}

	// Validating fields
	validate := validator.New()
	if err = validate.Struct(post); err != nil {
		return err
	}
	ha.post.NickName = post.NickName
	ha.post.Content = post.Content
	return nil
}
