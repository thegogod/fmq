package topics

import (
	"github.com/go-chi/chi/v5"
	"github.com/thegogod/fmq/server/storage"
)

func Router(topics *storage.Topics) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", Get(topics))
	r.Get("/{key}", GetOne(topics))
	return r
}
