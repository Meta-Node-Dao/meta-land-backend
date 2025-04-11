package image

// ListResponse
// tag list response
type ListResponse struct {
	List  []Image `json:"list"`
	Total int64   `json:"total"`
}
