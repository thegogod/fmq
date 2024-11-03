package admin

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Serve() http.HandlerFunc {
	fs := http.FileServer(http.Dir("../web/dist/web/browser"))

	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("../web/dist/web/browser", strings.TrimPrefix(r.URL.Path, "/admin"))
		fi, err := os.Stat(path)

		if os.IsNotExist(err) || fi.IsDir() {
			http.ServeFile(w, r, filepath.Join("../web/dist/web/browser", "index.html"))
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.StripPrefix("/admin", fs).ServeHTTP(w, r)
	}
}
