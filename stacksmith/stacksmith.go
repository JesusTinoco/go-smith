package stacksmith

import (
	"net/http"

	"github.com/dghubble/sling"
)

const stacksmithAPI = "https://stacksmith.bitnami.com/api/v1/"

// Client is a Stacksmith client for making Stacksmith API requests.
type Client struct {
	sling     *sling.Sling
	Stacks    *StacksService
	Hooks     *HooksService
	Discovery *DiscoveryService
	User      *UserService
}

// APIKeyParam ...
type APIKeyParam struct {
	APIKey string `url:"api_key"`
}

// NewClient return a new Client
func NewClient(apiKey string, httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(stacksmithAPI)
	return &Client{
		sling:     base,
		Stacks:    newStacksService(base.New().QueryStruct(APIKeyParam{APIKey: apiKey})),
		Hooks:     newHooksService(base.New().QueryStruct(APIKeyParam{APIKey: apiKey})),
		Discovery: newDiscoveryService(base.New().QueryStruct(APIKeyParam{APIKey: apiKey})),
		User:      newUserService(base.New().QueryStruct(APIKeyParam{APIKey: apiKey})),
	}
}
