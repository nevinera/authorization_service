package main

import (
  "os"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/nevinera/authorization_service/db"
  "github.com/nevinera/authorization_service/ctrl"
)

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
  router.Handle("/users/{uuid}", ctrl.UsersDestroyHandler(conn)).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":3000", router))
}
