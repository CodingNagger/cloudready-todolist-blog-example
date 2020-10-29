package creation

import (
	"fmt"
	"testing"
)

var spyTask Task

type HappyPathRepository struct {
}

func (r HappyPathRepository) CreateTask(t Task) error {
	spyTask = t
	return nil
}

func TestCreateTask_PassesTaskToRepositoryWithoutError(t *testing.T) {
	task := Task{Description: "test description"}
	repository := HappyPathRepository{}
	service := NewService(repository)

	err := service.CreateTask(task)

	if err != nil {
		t.Fatalf("An error was returned %s", err)
	}

	if spyTask != task {
		t.Fatalf("The task was not passed for creation %v - %v", spyTask, task)
	}
}

func TestCreateTask_FailsWhenDescriptionIsEmpty(t *testing.T) {
	task := Task{}
	repository := HappyPathRepository{}
	service := NewService(repository)

	err := service.CreateTask(task)

	if err != ErrorEmptyDescription {
		t.Fatalf("Did not fail with ErrorEmptyDescription")
	}

	if spyTask == task {
		t.Fatalf("The task was mistakenly passed for creation despite failed check")
	}
}

type ErroneousRepository struct {
	err error
}

func (r ErroneousRepository) CreateTask(t Task) error {
	return r.err
}

func TestCreateTask_FailsWithCorrectError(t *testing.T) {
	task := Task{Description: "test description"}
	repository := ErroneousRepository{fmt.Errorf("test error")}
	service := NewService(repository)

	err := service.CreateTask(task)

	if err != repository.err {
		t.Fatalf("Wrong error returned %s", err)
	}
}
