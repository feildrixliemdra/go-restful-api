package router

import (
	"github.com/feildrixliemdra/go-restful-api/internal/handler/http"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/").Subrouter()

	http.NewUserHandler(r)

	return r
}
