package stacksmith

// PaginationParams ...
type PaginationParams struct {
	Page    int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
}

// StatusGeneration ...
type StatusGeneration struct {
	ID       string `json:"id"`
	StackURL string `json:"stack_url"`
}

// StatusDeletion ...
type StatusDeletion struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}
