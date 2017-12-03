package ctrl

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/nevinera/authorization_service/data"
)

func UsersShowHandler(conn *data.Connection) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    user, found, err := conn.GetUser(params["uuid"])
    if err != nil {
      sendError(w, http.StatusInternalServerError, err.Error())
    } else if found != true {
      sendError(w, http.StatusNotFound, "User not found")
    } else {
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(user)
    }
  })
}

func UsersCreateHandler(conn *data.Connection) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    user, created, err := conn.CreateUser(params["uuid"])

    if err != nil {
      sendError(w, http.StatusInternalServerError, err.Error())
    } else if created == false {
      sendError(w, http.StatusNotAcceptable, "Unable to create user")
    } else {
      w.WriteHeader(http.StatusCreated)
      json.NewEncoder(w).Encode(user)
    }
  })
}

func UsersDestroyHandler(conn *data.Connection) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    user, destroyed, err := conn.DestroyUser(params["uuid"])

    if err != nil {
      sendError(w, http.StatusInternalServerError, err.Error())
    } else if user == nil {
      sendError(w, http.StatusNotFound, "User not found")
    } else if destroyed == false {
      sendError(w, http.StatusInternalServerError, "Unable to destroy user")
    } else {
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(user)
    }
  })
}
