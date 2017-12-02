package main

import (
  "log"
  "strconv"
  "net/http"
  "time"
  "encoding/json"
  "github.com/gorilla/mux"
)

type User struct {
  Uuid string
  CreatedAt time.Time
}

func sendError(w http.ResponseWriter, status int, msg string) {
  w.WriteHeader(status)
  errorData := map[string]string{
    "error": msg,
    "status": strconv.Itoa(status),
    "occurred_at": time.Now().String(),
  }
  json.NewEncoder(w).Encode(errorData)
}

var users = make(map[string]User)

func GetUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  if user, included := users[params["uuid"]]; included {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
  } else {
    sendError(w, http.StatusNotFound, "That user was not found")
  }
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  if _, included := users[params["uuid"]]; included {
    sendError(w, http.StatusConflict, "That user already exists")
  } else {
    user := User{Uuid: params["uuid"], CreatedAt: time.Now()}
    users[params["uuid"]] = user
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
  }
}

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
  router := mux.NewRouter()
  router.HandleFunc("/users/{uuid}", GetUser).Methods("GET")
  router.HandleFunc("/users/{uuid}", CreateUser).Methods("PUT")
  router.HandleFunc("/users/{uuid}", DestroyUser).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":3000", router))
}
