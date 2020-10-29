package listing

// ListedTasks is a collection of listed tasks
type ListedTasks []Task

// Service provides task listing operations.
type Service interface {
	ListCompletedTasks() (ListedTasks, error)
	ListPendingTasks() (ListedTasks, error)
	ListAllTasks() (ListedTasks, error)
}

// Repository provides access to the task repository.
type Repository interface {
	ListCompletedTasks() (ListedTasks, error)
	ListPendingTasks() (ListedTasks, error)
	ListAllTasks() (ListedTasks, error)
}

type service struct {
	tR Repository
}

// NewService creates a creation service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ListCompletedTasks() (ListedTasks, error) {
	return s.tR.ListCompletedTasks()
}

func (s *service) ListPendingTasks() (ListedTasks, error) {
	return s.tR.ListPendingTasks()
}

func (s *service) ListAllTasks() (ListedTasks, error) {
	return s.tR.ListAllTasks()
}
