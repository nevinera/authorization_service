package server

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/nevinera/authorization_service/data"
  "github.com/nevinera/authorization_service/ctrl"
)

func Run(listenString string, dbUrl string) error {
  conn, err := data.NewConnection(dbUrl)
  if err != nil {
    return err
  }

  conn.CreateDatabase()

  router := buildRouter(conn)
  http.ListenAndServe(listenString, router)
  return nil
}

func buildRouter(conn *data.Connection) *mux.Router {
  router := mux.NewRouter()

  router.Handle("/users/{uuid}", ctrl.UsersShowHandler(conn)).Methods("GET")
  router.Handle("/users/{uuid}", ctrl.UsersCreateHandler(conn)).Methods("PUT")
  router.Handle("/users/{uuid}", ctrl.UsersDestroyHandler(conn)).Methods("DELETE")

  router.Handle("/groups/{uuid}", ctrl.GroupsShowHandler(conn)).Methods("GET")
  router.Handle("/groups/{uuid}", ctrl.GroupsCreateHandler(conn)).Methods("PUT")
  router.Handle("/groups/{uuid}", ctrl.GroupsDestroyHandler(conn)).Methods("DELETE")

  return router
}
