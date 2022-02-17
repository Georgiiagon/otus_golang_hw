package internalhttp

import (
	"net/http"

	"github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/config"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func (apiHandler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func NewHandler(config.Config) (http.Handler, error) {
	return apiHandler{}, nil
}
