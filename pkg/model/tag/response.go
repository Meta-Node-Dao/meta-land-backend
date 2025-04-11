package tag

// ListResponse
// tag list response
type ListResponse struct {
	List  []Tag `json:"list"`
	Total int64 `json:"total"`
}
