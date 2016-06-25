package stacksmith

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// StacksService ...
type StacksService struct {
	sling *sling.Sling
}

func newStacksService(sling *sling.Sling) *StacksService {
	return &StacksService{
		sling: sling.Path("stacks/"),
	}
}

// StacksList ...
type StacksList struct {
	TotalEntries int     `json:"total_entries"`
	TotalPages   int     `json:"total_pages"`
	Items        []Stack `json:"items"`
}

// Stack ...
type Stack struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Status               string `json:"status"`
	GeneratedAt          string `json:"generated_at"`
	RegeneratedAt        string `json:"regenerated_at"`
	CanRegenerate        bool   `json:"can_regenerate"`
	Outdated             bool   `json:"outdated"`
	NotificationsEnabled bool   `json:"notifications_enabled"`
	Kind                 string `json:"kind"`
	Requirements         []struct {
		ID      string `json:"id"`
		Version string `json:"version"`
	} `json:"requirements"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
	Components      []Component     `json:"components"`
	Os              Component       `json:"os"`
	Output          struct {
		Dockerfile string `json:"dockerfile"`
	} `json:"output"`
	Shared       bool   `json:"shared"`
	ShareableURL string `json:"shareable_url"`
}

// Component ...
type Component struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Branch          string          `json:"branch"`
	Version         string          `json:"version"`
	Revision        int             `json:"revision"`
	Category        string          `json:"category"`
	Latest          Latest          `json:"latest"`
	Outdated        bool            `json:"outdated"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}

// Latest ...
type Latest struct {
	Version  string `json:"version"`
	Revision int    `json:"revision"`
}

// Vulnerability ...
type Vulnerability struct {
	Vulnerable   bool   `json:"vulnerable"`
	Severity     string `json:"severity"`
	URL          string `json:"url"`
	TotalEntries int    `json:"total_entries"`
	TotalPages   int    `json:"total_pages"`
	Items        []struct {
		Name     string `json:"name"`
		Severity string `json:"severity"`
		Ranges   []struct {
			Component string `json:"component"`
			From      string `json:"from"`
			To        string `json:"to"`
		} `json:"ranges"`
	} `json:"items"`
}

// StackDefinition ...
type StackDefinition struct {
	Name       string `json:"name"`
	Components []Item `json:"components"`
	OS         Item   `json:"os"`
	Kind       string `json:"kind"`
}

// Item ...
type Item struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

// StackParams ...
type StackParams struct {
	Name                 string `json:"name"`
	NotificationsEnabled bool   `json:"notifications_enabled"`
	Shared               bool   `json:"shared"`
}

// List ...
func (s *StacksService) List(params *PaginationParams) (*StacksList, *http.Response, error) {
	stacksList := new(StacksList)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(params).Receive(stacksList, apiError)
	return stacksList, resp, relevantError(err, *apiError)
}

// Create ...
func (s *StacksService) Create(params *StackDefinition) (*StatusGeneration, *http.Response, error) {
	status := new(StatusGeneration)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("").BodyJSON(params).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Delete ...
func (s *StacksService) Delete(stackID string) (*StatusDeletion, *http.Response, error) {
	status := new(StatusDeletion)
	apiError := new(APIError)
	resp, err := s.sling.New().Delete(stackID).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Get ...
func (s *StacksService) Get(stackID string) (*Stack, *http.Response, error) {
	stack := new(Stack)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(stackID).Receive(stack, apiError)
	return stack, resp, relevantError(err, *apiError)
}

// Update ...
func (s *StacksService) Update(stackID string, params *StackParams) (*StatusGeneration, *http.Response, error) {
	status := new(StatusGeneration)
	apiError := new(APIError)
	resp, err := s.sling.New().Patch(stackID).BodyJSON(params).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Regenerate ...
func (s *StacksService) Regenerate(stackID string) (*StatusGeneration, *http.Response, error) {
	status := new(StatusGeneration)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/regenerate", stackID)
	resp, err := s.sling.New().Post(path).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// GetVulnerabilities ...
func (s *StacksService) GetVulnerabilities(stackID string, params *PaginationParams) (*Vulnerability, *http.Response, error) {
	vulnerabilities := new(Vulnerability)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/vulnerabilities", stackID)
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(vulnerabilities, apiError)
	return vulnerabilities, resp, relevantError(err, *apiError)
}
