package stacksmith

import (
	"net/http"

	"github.com/dghubble/sling"
)

// DiscoveryService ...
type DiscoveryService struct {
	sling *sling.Sling
}

func newDiscoveryService(sling *sling.Sling) *DiscoveryService {
	return &DiscoveryService{
		sling: sling,
	}
}

// ListItem ...
type ListItem struct {
	Items []Item `json:"items"`
}

// Item ...
type Item struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Versions []struct {
		Version     string `json:"version"`
		Revision    int    `json:"revision"`
		Branch      string `json:"branch"`
		Checksum    string `json:"checksum"`
		PublishedAt string `json:"publised_at"`
	} `json:"versions"`
	Prebuilt      bool `json:"prebuilt"`
	ReleaseSeries []struct {
		Version string `json:"version"`
		Payload string `json:"payload"`
	} `json:"release_series"`
	DependenciesURL string `json:"dependencies_url"`
}

// ComponentsList ...
func (s *DiscoveryService) ComponentsList(query string) (*ListItem, *http.Response, error) {
	return getDiscovery(s, "components", query)
}

// ServicesList ...
func (s *DiscoveryService) ServicesList(query string) (*ListItem, *http.Response, error) {
	return getDiscovery(s, "services", query)
}

// RuntimesList ...
func (s *DiscoveryService) RuntimesList(query string) (*ListItem, *http.Response, error) {
	return getDiscovery(s, "runtimes", query)
}

// OsesList ...
func (s *DiscoveryService) OsesList(query string) (*ListItem, *http.Response, error) {
	return getDiscovery(s, "oses", query)
}

func getDiscovery(s *DiscoveryService, path string, query string) (*ListItem, *http.Response, error) {
	queryAux := struct{ query string }{query: query}
	componentList := new(ListItem)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(path).QueryStruct(queryAux).Receive(componentList, apiError)
	return componentList, resp, relevantError(err, *apiError)
}
