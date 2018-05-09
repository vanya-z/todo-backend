package task

import (
	"fmt"
	"net/http"
)

func AllTasksEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet AllTasksEndPoint!")
}

func FindTaskEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet FindTaskEndpoint!")
}

func CreateTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet CreateTaskEndPoint!")
}

func UpdateTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet UpdateTaskEndPoint!")
}

func DeleteTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet DeleteTaskEndPoint!")
}
