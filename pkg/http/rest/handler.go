package rest

import (
	"encoding/json"
	"net/http"

	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/completion"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/creation"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/listing"
	"github.com/julienschmidt/httprouter"
)

// Handler defines http endpoint
func Handler(cr creation.Service, l listing.Service, cp completion.Service) http.Handler {
	router := httprouter.New()

	router.GET("/tasks", listAllTasks(l))
	router.GET("/tasks/completed", listCompletedTasks(l))
	router.GET("/tasks/pending", listPendingTasks(l))

	router.POST("/tasks", createTask(cr))
	router.POST("/tasks/:id/complete", completeTask(cp))

	return router
}

// createTask returns a handler for POST /tasks requests
func createTask(s creation.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newTask creation.Task
		err := decoder.Decode(&newTask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.CreateTask(newTask)

		if err != nil {
			if err == creation.ErrorEmptyDescription {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// completeTask returns a handler for POST /tasks/:id/complete requests
func completeTask(s completion.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		taskID := p.ByName("id")
		taskToComplete := completion.Task{ID: taskID}

		err := s.CompleteTask(taskToComplete)

		if err != nil {
			if err == completion.ErrorTaskAlreadyCompleted {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else if err == completion.ErrorTaskNotFound {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

// listAllTasks returns a handler for GET /posts requests ; ignore errors for simplicity
func listAllTasks(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list, _ := s.ListAllTasks()
		json.NewEncoder(w).Encode(list)
	}
}

// listPendingTasks returns a handler for GET /posts requests ; ignore errors for simplicity
func listPendingTasks(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list, _ := s.ListPendingTasks()
		json.NewEncoder(w).Encode(list)
	}
}

// listCompletedTasks returns a handler for GET /posts requests ; ignore errors for simplicity
func listCompletedTasks(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list, _ := s.ListCompletedTasks()
		json.NewEncoder(w).Encode(list)
	}
}
