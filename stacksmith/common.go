package stacksmith

// PaginationParams ...
type PaginationParams struct {
	Page    int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
}

// StatusDeletion ...
type StatusDeletion struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}
