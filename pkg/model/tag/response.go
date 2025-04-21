package tag

// ListResponse
// tag list response
type ListResponse struct {
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
	Total int64 `json:"total"`
	List  []Tag `json:"list"`
}
