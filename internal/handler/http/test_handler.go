package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type userHandler struct {
}

func NewUserHandler(r *mux.Router) {
	handler := &userHandler{}
	v1 := r.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/test", handler.Test).Methods(http.MethodGet)
}

func (h *userHandler) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("success %s", r.URL)))
}
