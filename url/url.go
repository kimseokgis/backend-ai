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
	case method == "GET" && path == "/getuser":
		controller.GetUser(w, r)
	case method == "GET" && path == "/getallusers":
		controller.GetAllUsers(w, r)
	case method == "POST" && path == "/chatDomykado":
		controller.ChatPredictForDomykado(w, r)
	case method == "POST" && path == "/chatRegexp":
		controller.ChatPredictUsingRegexp(w, r)
	case method == "POST" && path == "/comment":
		controller.Comment(w, r)
	default:
		controller.NotFound(w, r)
	}
}
