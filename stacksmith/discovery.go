package stacksmith

import (
	"fmt"
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

// Kinds ...
type Kinds struct {
	TotalEntries int `json:"total_entries"`
	TotalPages   int `json:"total_pages"`
	Items        []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Default      bool   `json:"default"`
		ComponentURL string `json:"component_url"`
	} `json:"items"`
}

// Changelog ...
type Changelog struct {
	TotalEntries int `json:"total_entries"`
	TotalPAges   int `json:"total_pages"`
	Items        []struct {
		Version         string `json:"version"`
		Revision        int    `json:"revision"`
		Branch          string `json:"branch"`
		Checksum        string `json:"checksum"`
		PublishedAt     string `json:"published_at"`
		ReleaseNotes    string `json:"release_notes"`
		ReleaseNotesURL string `json:"release_notes_url"`
	} `json:"items"`
}

// Dependencies ...
type Dependencies struct {
	TotalEntries int      `json:"total_entries"`
	TotalPages   int      `json:"total_pages"`
	Items        []string `json:"items"`
}

// Query ...
type Query struct {
	Query string `url:"query,omitempty"`
}

// RangeParams ...
type RangeParams struct {
	From string `url:"from,omitempty"`
	To   string `url:"to,omitempty"`
}

// ComponentsList ...
func (s *DiscoveryService) ComponentsList(query string) (*ListItem, *http.Response, error) {
	return getDiscovery(s, "components", query)
}

// GetComponent ...
func (s *DiscoveryService) GetComponent(componentName string) (*Item, *http.Response, error) {
	component := new(Item)
	apiError := new(APIError)
	path := fmt.Sprintf("components/%s", componentName)
	resp, err := s.sling.New().Get(path).Receive(component, apiError)
	return component, resp, relevantError(err, *apiError)
}

// GetChangelogFrom ...
func (s *DiscoveryService) GetChangelogFrom(componentName string,
	rangeParam *RangeParams, pageParam *PaginationParams) (*Changelog, *http.Response, error) {
	changelog := new(Changelog)
	apiError := new(APIError)
	path := fmt.Sprintf("components/%s/changelog", componentName)
	resp, err := s.sling.New().Get(path).QueryStruct(rangeParam).QueryStruct(pageParam).Receive(changelog, apiError)
	return changelog, resp, relevantError(err, *apiError)
}

// GetDependenciesFrom ...
func (s *DiscoveryService) GetDependenciesFrom(componentName string) (*Dependencies, *http.Response, error) {
	dependencies := new(Dependencies)
	apiError := new(APIError)
	path := fmt.Sprintf("components/%s/dependencies", componentName)
	resp, err := s.sling.New().Get(path).Receive(dependencies, apiError)
	return dependencies, resp, relevantError(err, *apiError)
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

// KindsList ...
func (s *DiscoveryService) KindsList(pageParams *PaginationParams) (*Kinds, *http.Response, error) {
	kinds := new(Kinds)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("kinds").QueryStruct(pageParams).Receive(kinds, apiError)
	return kinds, resp, relevantError(err, *apiError)
}

// GetKindsFrom ...
func (s *DiscoveryService) GetKindsFrom(componentName string,
	pageParams *PaginationParams) (*Kinds, *http.Response, error) {
	kinds := new(Kinds)
	apiError := new(APIError)
	path := fmt.Sprintf("components/%s/kinds", componentName)
	resp, err := s.sling.New().Get(path).Receive(kinds, apiError)
	return kinds, resp, relevantError(err, *apiError)
}

func getDiscovery(s *DiscoveryService, path string, query string) (*ListItem, *http.Response, error) {
	componentList := new(ListItem)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(path).QueryStruct(Query{Query: query}).Receive(componentList, apiError)
	return componentList, resp, relevantError(err, *apiError)
}
