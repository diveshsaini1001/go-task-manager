package task

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
	IsCompleted bool   `json:"is_completed"`
}
