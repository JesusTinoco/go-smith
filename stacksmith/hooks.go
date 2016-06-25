package stacksmith

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// HooksService ...
type HooksService struct {
	sling *sling.Sling
}

func newHooksService(sling *sling.Sling) *HooksService {
	return &HooksService{
		sling: sling.Path("stacks/"),
	}
}

// HooksList ...
type HooksList struct {
	TotalEntries int `json:"total_entries"`
	TotalPages   int `json:"total_pages"`
	Items        []struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"items"`
}

// TestHook ...
type TestHook struct {
	ID     string `json:"id"`
	Result struct {
		Request struct {
			URL  string `json:"url"`
			Body string `json:"body"`
		} `json:"request"`
	} `json:"result"`
	Response struct {
		Code    string `json:"code"`
		Body    string `json:"body"`
		Message string `json:"message"`
	}
}

// ResponseGeneration ...
type ResponseGeneration struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// HookParams ...
type HookParams struct {
	URL string `json:"url"`
}

// List ...
func (s *HooksService) List(stackID string, params *PaginationParams) (*HooksList, *http.Response, error) {
	hooksList := new(HooksList)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/hooks", stackID)
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(hooksList, apiError)
	return hooksList, resp, relevantError(err, *apiError)
}

// Register ...
func (s *HooksService) Register(stackID string, params *HookParams) (*ResponseGeneration, *http.Response, error) {
	status := new(ResponseGeneration)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/hooks", stackID)
	resp, err := s.sling.New().Post(path).BodyJSON(params).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Delete ...
func (s *HooksService) Delete(stackID string, hookID string) (*StatusDeletion, *http.Response, error) {
	status := new(StatusDeletion)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/hooks/%s", stackID, hookID)
	resp, err := s.sling.New().Delete(path).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Update ...
func (s *HooksService) Update(stackID string, hookID string, params *HookParams) (*ResponseGeneration, *http.Response, error) {
	status := new(ResponseGeneration)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/hooks/%s", stackID, hookID)
	resp, err := s.sling.New().Patch(path).BodyJSON(params).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Test ...
func (s *HooksService) Test(stackID string, hookID string) (*TestHook, *http.Response, error) {
	testHook := new(TestHook)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/hooks/%s/test", stackID, hookID)
	resp, err := s.sling.New().Post(path).Receive(testHook, apiError)
	return testHook, resp, relevantError(err, *apiError)
}
