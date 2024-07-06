package main

import (
	"github.com/kimseokgis/backend-ai/url"
	"net/http"
)

func main() {
	http.HandleFunc("/", url.Web)
	http.ListenAndServe(":8080", nil)
}
