package admin

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	fs := http.FileServer(http.Dir("../web/dist/web/browser"))
	r := chi.NewRouter()
	r.Handle("/", http.StripPrefix("/admin", fs))
	r.HandleFunc("/*", Serve())

	return r
}
