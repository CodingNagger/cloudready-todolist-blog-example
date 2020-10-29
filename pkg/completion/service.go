package completion

import "errors"

var (
	// ErrorTaskAlreadyCompleted is an error returned when attempting to complete an already completed task
	ErrorTaskAlreadyCompleted = errors.New("Task already completed")
	// ErrorTaskNotFound is an error returned when a task is not found
	ErrorTaskNotFound = errors.New("Task not found")
)

// Service provides task completion operations.
type Service interface {
	CompleteTask(Task) error
}

// Repository provides access to the task repository.
type Repository interface {
	CompleteTask(Task) error
}

type service struct {
	tR Repository
}

// NewService creates a creation service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// CreateTasks creates the given tasks
func (s *service) CompleteTask(t Task) error {
	return s.tR.CompleteTask(t)
}
