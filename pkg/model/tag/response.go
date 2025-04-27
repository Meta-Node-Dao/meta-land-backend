package tag

// ListResponse
// tag list response
type ListResponse struct {
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
	Total int64 `json:"total"`
	List  []Tag `json:"list"`
}
type TagRelationResponse struct {
	Id       int `json:"id"`
	Tag      TagResponse
	TagId    int `json:"tag_id"`
	TargetId int `json:"target_id"`
	Type     int `json:"type"`
}

type TagResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}
