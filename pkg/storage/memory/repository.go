package memory

import (
	"github.com/google/uuid"

	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/completion"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/creation"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/listing"
)

// Storage keeps tasks data in memory
type Storage struct {
	tasks []*Task
}

// NewStorage creates and return a new storage
func NewStorage() *Storage {
	return new(Storage)
}

// CreateTask adds a new task in memory
func (s *Storage) CreateTask(t creation.Task) error {
	newTask := Task{
		ID:          uuid.New().String(),
		Description: t.Description,
		Completed:   false,
	}

	s.tasks = append(s.tasks, &newTask)

	return nil
}

// CompleteTask marks a task as completed
func (s *Storage) CompleteTask(task completion.Task) error {
	for _, t := range s.tasks {
		if task.ID == t.ID {
			if t.Completed {
				return completion.ErrorTaskAlreadyCompleted
			}

			t.Completed = true
			return nil
		}
	}

	return completion.ErrorTaskNotFound
}

// ListCompletedTasks lists all the tasks marked as completed
func (s *Storage) ListCompletedTasks() (listing.ListedTasks, error) {
	result := listing.ListedTasks{}

	for _, t := range s.tasks {
		if !t.Completed {
			continue
		}

		task := listing.Task{
			ID:          t.ID,
			Description: t.Description,
		}

		result = append(result, task)
	}

	return result, nil
}

// ListPendingTasks lists all the tasks not marked as completed
func (s *Storage) ListPendingTasks() (listing.ListedTasks, error) {
	result := listing.ListedTasks{}

	for _, t := range s.tasks {
		if t.Completed {
			continue
		}

		task := listing.Task{
			ID:          t.ID,
			Description: t.Description,
		}

		result = append(result, task)
	}

	return result, nil
}

// ListAllTasks returns all tasks
func (s *Storage) ListAllTasks() (listing.ListedTasks, error) {
	result := listing.ListedTasks{}

	for _, t := range s.tasks {
		task := listing.Task{
			ID:          t.ID,
			Description: t.Description,
		}

		result = append(result, task)
	}

	return result, nil
}
