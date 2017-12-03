package main

import (
  "os"
  "log"
  "net/http"
  "time"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/nevinera/authorization_service/db"
  "github.com/nevinera/authorization_service/ctrl"
)

type User struct {
  Uuid string
  CreatedAt time.Time
}

type httpError struct {
  message string `json:"error"`
  status int
  occurred_at time.Time
}

func sendError(w http.ResponseWriter, status int, msg string) {
  w.WriteHeader(status)
  errorData := httpError{message: msg, status: status, occurred_at: time.Now()}
  json.NewEncoder(w).Encode(errorData)
}

var users = make(map[string]User)

func DestroyUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  if user, included := users[params["uuid"]]; included {
    delete(users, params["uuid"])
    json.NewEncoder(w).Encode(user)
  } else {
    sendError(w, http.StatusNotFound, "That user was not found")
  }
}

func main() {
  connectionString, wasSet := os.LookupEnv("DATABASE_URL")
  if !wasSet {
    log.Fatal("DATABASE_URL must be set")
  }

  conn, err := db.NewConnection(connectionString)
  if err != nil {
    log.Fatal("Could not connect to database")
  }

  conn.CreateDatabase()

  router := mux.NewRouter()
  router.Handle("/users/{uuid}", ctrl.UsersShowHandler(conn)).Methods("GET")
  router.Handle("/users/{uuid}", ctrl.UsersCreateHandler(conn)).Methods("PUT")
  router.HandleFunc("/users/{uuid}", DestroyUser).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":3000", router))
}
