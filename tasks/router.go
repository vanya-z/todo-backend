package tasks

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

// Route defines a route
type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
  Route{
    "List of tasks",
    "GET",
    "/api/v1/tasks",
    AllTasksEndPoint,
  },
  Route{
    "Create a task",
    "POST",
    "/api/v1/tasks",
    CreateTaskEndPoint,
  },
  Route{
    "Toggle the task",
    "PUT",
    "/api/v1/tasks/{id}/toggle",
    ToggleTaskEndPoint,
  },
  Route{
    "Delete the task",
    "DELETE",
    "/api/v1/tasks/{id}",
    DeleteTaskEndPoint,
  },
}

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)

  for _, route := range routes {
    var handler http.Handler
    log.Println(route.Name)
    handler = route.HandlerFunc

    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)
  }
  return router
}
