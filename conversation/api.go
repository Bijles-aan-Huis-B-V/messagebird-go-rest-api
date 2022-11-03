package conversation

import (
	"fmt"
	messagebird "github.com/Bijles-aan-Huis-B-V/messagebird-go-rest-api"
)

const (
	// apiRoot is the absolute URL of the Converstations API. All paths are
	// relative to apiRoot (e.g.
	// https://conversations.messagebird.com/v1/webhooks).
	apiRoot = "https://conversations.messagebird.com/v1"

	// path is the path for the Conversation resource, relative to apiRoot.
	path = "conversations"

	// startConversationPath is the path for starting new conversation
	startConversationPath = "start"

	// contactPath is the path for fetching a collection of conversations by contact ID
	contactPath = "contact"

	// messagesPath is the path for the Message resource, relative to apiRoot
	// and path.
	messagesPath = "messages"

	// sendMessagePath is the path for creating the Message resource relative to apiRoot
	sendMessagePath = "send"

	// webhooksPath is the path for the Webhook resource, relative to apiRoot.
	webhooksPath = "webhooks"
)

// request does the exact same thing as DefaultClient.Request. It does, however,
// prefix the path with the Conversation API's root. This ensures the client
// doesn't "handle" this for us: by default, it uses the REST API.
func request(c messagebird.Client, v interface{}, method, path string, data interface{}) error {
	return c.Request(v, method, fmt.Sprintf("%s/%s", apiRoot, path), data)
}
