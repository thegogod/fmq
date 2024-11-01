package topics

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/thegogod/fmq/server/storage"
)

func Get(topics *storage.Topics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, topics)
	}
}
