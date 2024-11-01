package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/thegogod/fmq/server/routes/topics"
	"github.com/thegogod/fmq/server/storage"
)

func Router(state *storage.Topics) *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/topics", topics.Router(state))
	return r
}
