package handler

import (
	"encoding/json"
	"net/http"

	"github.com/danya1off/library/service"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	libService service.Library
}

func New(libService service.Library) *Handlers {
	return &Handlers{libService: libService}
}

func (h *Handlers) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.libService.GetAllBooks(r.Context())
	if err != nil {
		logrus.WithError(err).Errorf("Failed to load books")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(books); err != nil {
		logrus.WithError(err).Errorf("Failed encode books to json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
