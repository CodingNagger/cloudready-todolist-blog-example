package creation

import "errors"

var (
	// ErrorEmptyDescription is an error returned when attempting to create a task with an empty description
	ErrorEmptyDescription = errors.New("Description cannot be empty")
)

// Service provides task creation operations.
type Service interface {
	CreateTask(Task) error
}

// Repository provides access to the task repository.
type Repository interface {
	CreateTask(Task) error
}

type service struct {
	tR Repository
}

// NewService creates a creation service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// CreateTasks creates the given tasks
func (s *service) CreateTask(t Task) error {
	if len(t.Description) == 0 {
		return ErrorEmptyDescription
	}

	return s.tR.CreateTask(t)
}
