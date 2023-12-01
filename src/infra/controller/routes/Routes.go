package routes

import (
	"io"
	"net/http"

	"github.com/isaquecsilva/posts-rest-api/core/usecase"
	httpAdapter "github.com/isaquecsilva/posts-rest-api/infra/adapters/http"
	"github.com/isaquecsilva/posts-rest-api/infra/adapters/repository/inmem"
	"github.com/isaquecsilva/posts-rest-api/infra/controller"
)

var (
	repo    = &inmem.RepositoryMock{}
	service = usecase.NewPostService(repo)
	control = controller.NewController(service)
)

func InitRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Write([]byte("Not found"))
			return
		}

		var adapter = httpAdapter.NewHttpAdapter(r, w)
		control.CreateNewPost(adapter)
	})

	http.HandleFunc("/find", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Write([]byte("Not found"))
			return
		}

		var adapter = httpAdapter.NewHttpAdapter(r, w)
		control.GetPost(adapter)
	})

	http.HandleFunc("/findall", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Write([]byte("Not found"))
			return
		}

		var adapter = httpAdapter.NewHttpAdapter(r, w)
		control.GetAllPosts(adapter)
	})
}

func StartServer(addr string) error {
	return http.ListenAndServe(addr, nil)
}
