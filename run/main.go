package main

import (
	"net/http"

	"github.com/kimseokgis/backend-ai/url"
)

func main() {
	http.HandleFunc("/", url.Web)
	
	http.ListenAndServe(":8080", nil)
}
