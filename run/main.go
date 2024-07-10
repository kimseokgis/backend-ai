package main

import (
	"net/http"

	"github.com/kimseokgis/backend-ai/controller"
	"github.com/kimseokgis/backend-ai/url"
)

func main() {
	http.HandleFunc("/", url.Web)

	http.HandleFunc("/getuser", controller.GetUser)
	
	http.ListenAndServe(":8080", nil)
}
