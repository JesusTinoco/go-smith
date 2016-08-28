package stacksmith

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// DiscoveryService provides methods for accessing Stacksmith Discovery API endpoints.
type DiscoveryService struct {
	sling *sling.Sling
}

func newDiscoveryService(sling *sling.Sling) *DiscoveryService {
	return &DiscoveryService{
		sling: sling,
	}
}

// ListItems ...
type ListItems struct {
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
		PublishedAt string `json:"published_at"`
	} `json:"versions"`
	Prebuilt      bool `json:"prebuilt"`
	ReleaseSeries []struct {
		Version string `json:"version"`
		Payload string `json:"payload"`
	} `json:"release_series"`
	DependenciesURL string `json:"dependencies_url"`
}

// Flavors ...
type Flavors struct {
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

// ComponentsList List all available components.
// https://stacksmith.bitnami.com/api/v1/#!/Discovery/get_components
func (s *DiscoveryService) ComponentsList(query string) (*ListItems, *http.Response, error) {
	return getDiscovery(s, "components", query)
}

// GetComponent Retrieve the properties from a components
// https://stacksmith.bitnami.com/api/v1/#!/Discovery/get_components_id
func (s *DiscoveryService) GetComponent(componentName string) (*Item, *http.Response, error) {
	component := new(Item)
	apiError := new(APIError)
	path := fmt.Sprintf("components/%s", componentName)
	resp, err := s.sling.New().Get(path).Receive(component, apiError)
	return component, resp, relevantError(err, *apiError)
}

// GetChangelogFrom Retrieve the changelog for a component
// https://stacksmith.bitnami.com/api/v1/#!/Discovery/get_components_id_changelog
func (s *DiscoveryService) GetChangelogFrom(componentName string,
	rangeParam *RangeParams, pageParam *PaginationParams) (*Changelog, *http.Response, error) {
	changelog := new(Changelog)
	apiError := new(APIError)
	path := fmt.Sprintf("components/%s/changelog", componentName)
	resp, err := s.sling.New().Get(path).QueryStruct(rangeParam).QueryStruct(pageParam).Receive(changelog, apiError)
	return changelog, resp, relevantError(err, *apiError)
}

// GetDependenciesFrom Retrieve the component ID of the component dependencies
// https://stacksmith.bitnami.com/api/v1/#!/Discovery/get_components_id_dependencies
func (s *DiscoveryService) GetDependenciesFrom(componentName string) (*Dependencies, *http.Response, error) {
	dependencies := new(Dependencies)
	apiError := new(APIError)
	path := fmt.Sprintf("components/%s/dependencies", componentName)
	resp, err := s.sling.New().Get(path).Receive(dependencies, apiError)
	return dependencies, resp, relevantError(err, *apiError)
}

// ServicesList List all available components in the services category.
// https://stacksmith.bitnami.com/api/v1/#!/Discovery/get_services
func (s *DiscoveryService) ServicesList(query string) (*ListItems, *http.Response, error) {
	return getDiscovery(s, "services", query)
}

// RuntimesList List all available components in the runtimes category.
// https://stacksmith.bitnami.com/api/v1/#!/Discovery/get_runtimes
func (s *DiscoveryService) RuntimesList(query string) (*ListItems, *http.Response, error) {
	return getDiscovery(s, "runtimes", query)
}

// OsesList List all available OSes.
// https://stacksmith.bitnami.com/api/v1/#!/Discovery/get_oses
func (s *DiscoveryService) OsesList(query string) (*ListItems, *http.Response, error) {
	return getDiscovery(s, "oses", query)
}

// FlavorsList List all available Flavors
// https://stacksmith.bitnami.com/api/v1/#!/Discovery/get_flavors
func (s *DiscoveryService) FlavorsList(pageParams *PaginationParams) (*Flavors, *http.Response, error) {
	flavors := new(Flavors)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("flavors").QueryStruct(pageParams).Receive(flavors, apiError)
	return flavors, resp, relevantError(err, *apiError)
}

// GetFlavorsFrom Retrieve the available kinds from a component
// https://stacksmith.bitnami.com/api/v1/#!/Discovery/get_components_id_flavors
func (s *DiscoveryService) GetFlavorsFrom(componentName string,
	pageParams *PaginationParams) (*Flavors, *http.Response, error) {
	flavors := new(Flavors)
	apiError := new(APIError)
	path := fmt.Sprintf("components/%s/flavors", componentName)
	resp, err := s.sling.New().Get(path).QueryStruct(pageParams).Receive(flavors, apiError)
	return flavors, resp, relevantError(err, *apiError)
}

func getDiscovery(s *DiscoveryService, path string, query string) (*ListItems, *http.Response, error) {
	componentList := new(ListItems)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(path).QueryStruct(Query{Query: query}).Receive(componentList, apiError)
	return componentList, resp, relevantError(err, *apiError)
}
