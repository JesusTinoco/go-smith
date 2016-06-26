package stacksmith

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// StacksService provides methods for accessing Stacksmith Stacks API endpoints.
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

// StatusGeneration ...
type StatusGeneration struct {
	ID       string `json:"id"`
	StackURL string `json:"stack_url"`
}

// StackDefinition ...
type StackDefinition struct {
	Name       string          `json:"name"`
	Components []ComponentItem `json:"components"`
	OS         ComponentItem   `json:"os"`
	Kind       string          `json:"kind"`
}

// ComponentItem ...
type ComponentItem struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

// StackParams ...
type StackParams struct {
	Name                 string `json:"name"`
	NotificationsEnabled bool   `json:"notifications_enabled"`
	Shared               bool   `json:"shared"`
}

// List List all stacks attached to your account.
// https://stacksmith.bitnami.com/api/v1/#!/Stacks/get_stacks
func (s *StacksService) List(params *PaginationParams) (*StacksList, *http.Response, error) {
	stacksList := new(StacksList)
	apiError := new(APIError)
	resp, err := s.sling.New().QueryStruct(params).Receive(stacksList, apiError)
	return stacksList, resp, relevantError(err, *apiError)
}

// Create Create a stack by specifying the components you need, its kind and its OS.
// https://stacksmith.bitnami.com/api/v1/#!/Stacks/post_stacks
func (s *StacksService) Create(params *StackDefinition) (*StatusGeneration, *http.Response, error) {
	status := new(StatusGeneration)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("").BodyJSON(params).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Delete Delete a stack.
// https://stacksmith.bitnami.com/api/v1/#!/Stacks/delete_stacks_id
func (s *StacksService) Delete(stackID string) (*StatusDeletion, *http.Response, error) {
	status := new(StatusDeletion)
	apiError := new(APIError)
	resp, err := s.sling.New().Delete(stackID).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Get Retrieve the properties of a stack, to list the versions of the framework, runtime, and OS generated.
// https://stacksmith.bitnami.com/api/v1/#!/Stacks/get_stacks_id
func (s *StacksService) Get(stackID string) (*Stack, *http.Response, error) {
	stack := new(Stack)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(stackID).Receive(stack, apiError)
	return stack, resp, relevantError(err, *apiError)
}

// Update Update the properties of an existing stack.
// https://stacksmith.bitnami.com/api/v1/#!/Stacks/patch_stacks_id
func (s *StacksService) Update(stackID string, params *StackParams) (*StatusGeneration, *http.Response, error) {
	status := new(StatusGeneration)
	apiError := new(APIError)
	resp, err := s.sling.New().Patch(stackID).BodyJSON(params).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Regenerate Create a new stack based on the requirements of another, if there are new versions for it's requirements.
// https://stacksmith.bitnami.com/api/v1/#!/Stacks/post_stacks_id_regenerate
func (s *StacksService) Regenerate(stackID string) (*StatusGeneration, *http.Response, error) {
	status := new(StatusGeneration)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/regenerate", stackID)
	resp, err := s.sling.New().Post(path).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// GetVulnerabilities Retrieve the list of vulnerabilities affecting a stack.
// https://stacksmith.bitnami.com/api/v1/#!/Stacks/get_stacks_id_vulnerabilities
func (s *StacksService) GetVulnerabilities(stackID string, params *PaginationParams) (*Vulnerability, *http.Response, error) {
	vulnerabilities := new(Vulnerability)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/vulnerabilities", stackID)
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(vulnerabilities, apiError)
	return vulnerabilities, resp, relevantError(err, *apiError)
}
