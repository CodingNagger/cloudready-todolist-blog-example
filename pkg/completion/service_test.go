package completion

import (
	"testing"
)

type HappyPathRepository struct {
	t Task
}

func (r HappyPathRepository) CompleteTask(t Task) error {
	r.t = t
	return nil
}

func TestCompleteTask_PassesTaskToRepositoryWithoutError(t *testing.T) {
	task := Task{}
	repository := HappyPathRepository{}
	service := NewService(repository)

	err := service.CompleteTask(task)

	if err != nil {
		t.Fatalf("An error was returned %s", err)
	}

	if repository.t != task {
		t.Fatalf("The task was not completed")
	}
}

type TaskAlreadyCompletedRepository struct {
}

func (r TaskAlreadyCompletedRepository) CompleteTask(t Task) error {
	return ErrorTaskAlreadyCompleted
}

func TestCompleteTask_ShouldFailWithTaskAlreadyCompletedWhenDue(t *testing.T) {
	task := Task{}
	repository := TaskAlreadyCompletedRepository{}
	service := NewService(repository)

	err := service.CompleteTask(task)

	if err != ErrorTaskAlreadyCompleted {
		t.Fatalf("The wrong error error was returned %s", err)
	}
}

type TaskNotFoundRepository struct {
}

func (r TaskNotFoundRepository) CompleteTask(t Task) error {
	return ErrorTaskNotFound
}

func TestCompleteTask_ShouldFailWithTaskNotFoundWhenDue(t *testing.T) {
	task := Task{}
	repository := TaskNotFoundRepository{}
	service := NewService(repository)

	err := service.CompleteTask(task)

	if err != ErrorTaskNotFound {
		t.Fatalf("The wrong error error was returned %s", err)
	}
}
