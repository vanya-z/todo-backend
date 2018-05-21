package tasks

import (
  "encoding/json"
  "gopkg.in/mgo.v2/bson"
  "net/http"
  "github.com/gorilla/mux"
)

// GET list of tasks
func AllTasksEndPoint(w http.ResponseWriter, r *http.Request) {
  tasks, err := dbConfig.FindAll()
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, tasks)
}

// POST create a task
func CreateTaskEndPoint(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  var task Task
  if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    return
  }
  if task.Title == "" {
    respondWithError(w, http.StatusBadRequest, "Task cannot be empty")
    return
  }
  task.ID = bson.NewObjectId()
  if err := dbConfig.Insert(task); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusCreated, task)
}

// PUT toggle the task
func ToggleTaskEndPoint(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  task, err := dbConfig.FindById(params["id"])
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid Task ID")
    return
  }
  task.Completed = !task.Completed
  if err := dbConfig.Update(task); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, task)
}

// DELETE remove the task
func DeleteTaskEndPoint(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  task, err := dbConfig.FindById(params["id"])
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid Task ID")
    return
  }
  if err := dbConfig.Delete(task); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, task)
}


func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
  respondWithJson(w, code, map[string]string{"error": msg})
}
