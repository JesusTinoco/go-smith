package stacksmith

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// HooksService provides methods for accessing Stacksmith Stack Hooks API endpoints.
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

// List List all hooks for this stack.
// https://stacksmith.bitnami.com/api/v1/#!/Stack_Hooks/get_stacks_stack_id_hooks
func (s *HooksService) List(stackID string, params *PaginationParams) (*HooksList, *http.Response, error) {
	hooksList := new(HooksList)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/hooks", stackID)
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(hooksList, apiError)
	return hooksList, resp, relevantError(err, *apiError)
}

// Register Register a URL as a hook that will be triggered when there are updates for your stacks.
// https://stacksmith.bitnami.com/api/v1/#!/Stack_Hooks/post_stacks_stack_id_hooks
func (s *HooksService) Register(stackID string, params *HookParams) (*ResponseGeneration, *http.Response, error) {
	status := new(ResponseGeneration)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/hooks", stackID)
	resp, err := s.sling.New().Post(path).BodyJSON(params).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Delete Delete a hook
// https://stacksmith.bitnami.com/api/v1/#!/Stack_Hooks/delete_stacks_stack_id_hooks_id
func (s *HooksService) Delete(stackID string, hookID string) (*StatusDeletion, *http.Response, error) {
	status := new(StatusDeletion)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/hooks/%s", stackID, hookID)
	resp, err := s.sling.New().Delete(path).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Update Update the URL for a previously registered hook.
// https://stacksmith.bitnami.com/api/v1/#!/Stack_Hooks/patch_stacks_stack_id_hooks_id
func (s *HooksService) Update(stackID string, hookID string, params *HookParams) (*ResponseGeneration, *http.Response, error) {
	status := new(ResponseGeneration)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/hooks/%s", stackID, hookID)
	resp, err := s.sling.New().Patch(path).BodyJSON(params).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// Test Send a test payload to the URL endpoint.
// https://stacksmith.bitnami.com/api/v1/#!/Stack_Hooks/post_stacks_stack_id_hooks_id_test
func (s *HooksService) Test(stackID string, hookID string) (*TestHook, *http.Response, error) {
	testHook := new(TestHook)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/hooks/%s/test", stackID, hookID)
	resp, err := s.sling.New().Post(path).Receive(testHook, apiError)
	return testHook, resp, relevantError(err, *apiError)
}
