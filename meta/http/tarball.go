package http

import (
	"github.com/gaojiasheng/agent-updater/meta/g"
	"net/http"
)

func configTarballRoutes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(g.Config().TarballDir)).ServeHTTP(w, r)
	})

}
