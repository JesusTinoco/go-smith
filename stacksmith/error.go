package stacksmith

import "fmt"

// APIError ...
type APIError struct {
	Status  string `json:"status"`
	Message string `json:"error"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("stacksmith: %v %v", e.Status, e.Message)
}

// Empty ...
func (e APIError) Empty() bool {
	return false
}

// relevantError returns any non-nil http-related error (creating the request,
func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}
	return apiError
}
