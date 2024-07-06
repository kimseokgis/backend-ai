package backend_ai

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/kimseokgis/backend-ai/url"
)

func init() {
	functions.HTTP("makmur", url.Web)
}
