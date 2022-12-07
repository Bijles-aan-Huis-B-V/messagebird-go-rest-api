package integration

import (
	"fmt"
	messagebird "github.com/Bijles-aan-Huis-B-V/messagebird-go-rest-api"
)

const (
	// apiRoot is the absolute URL of the Integrations API. All paths are
	// relative to apiRoot.
	apiRoot = "https://integrations.messagebird.com"

	version = "v2"
	path    = "platforms"

	// whatsAppTemplatePath is the path for managing WhatsApp templates, relative to apiRoot.
	whatsAppTemplatePath = "whatsapp/templates"
)

// request does the exact same thing as DefaultClient.Request. It does, however,
// prefix the path with the Conversation API's root. This ensures the client
// doesn't "handle" this for us: by default, it uses the REST API.
func request(c messagebird.Client, v interface{}, method, path string, data interface{}) error {
	return c.Request(v, method, fmt.Sprintf("%s/%s", apiRoot, path), data)
}
