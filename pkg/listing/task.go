package listing

// A Task that can be listed
type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}
