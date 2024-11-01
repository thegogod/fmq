package topics

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/thegogod/fmq/server/storage"
)

func GetOne(topics *storage.Topics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		topic, exists := topics.Get(chi.URLParam(r, "key"))

		if !exists {
			render.Status(r, 404)
			render.JSON(w, r, "not found")
			return
		}

		render.JSON(w, r, topic)
	}
}
