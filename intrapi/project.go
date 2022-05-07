package intrapi

type Projects []Project
type Project struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Slug     string      `json:"slug"`
	ParentID interface{} `json:"parent_id"`
}
