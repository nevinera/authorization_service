package ctrl

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/nevinera/authorization_service/db"
)

func GroupsShowHandler(conn *db.Connection) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    group, found, err := conn.GetGroup(params["uuid"])
    if err != nil {
      sendError(w, http.StatusInternalServerError, err.Error())
    } else if found != true {
      sendError(w, http.StatusNotFound, "Group not found")
    } else {
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(group)
    }
  })
}

func GroupsCreateHandler(conn *db.Connection) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    group, created, err := conn.CreateGroup(params["uuid"])

    if err != nil {
      sendError(w, http.StatusInternalServerError, err.Error())
    } else if created == false {
      sendError(w, http.StatusNotAcceptable, "Unable to create group")
    } else {
      w.WriteHeader(http.StatusCreated)
      json.NewEncoder(w).Encode(group)
    }
  })
}

func GroupsDestroyHandler(conn *db.Connection) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    group, destroyed, err := conn.DestroyGroup(params["uuid"])

    if err != nil {
      sendError(w, http.StatusInternalServerError, err.Error())
    } else if group == nil {
      sendError(w, http.StatusNotFound, "Group not found")
    } else if destroyed == false {
      sendError(w, http.StatusInternalServerError, "Unable to destroy group")
    } else {
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(group)
    }
  })
}
