package url

import (
	"github.com/kimseokgis/backend-ai/config"
	"github.com/kimseokgis/backend-ai/controller"
	"net/http"
)

func Web(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return // If it's a preflight request, return early.
	}
	var method, path string = r.Method, r.URL.Path
	switch {
	case method == "GET" && path == "/":
		controller.HomeMakmur(w, r)
	case method == "POST" && path == "/registers":
		controller.RegisterUsers(w, r)
	case method == "POST" && path == "/login":
		controller.LoginUsers(w, r)
	default:
		controller.NotFound(w, r)
	}
}
