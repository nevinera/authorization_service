package handlers

import (
  "fmt"
  "net/http"
  "strings"
  "regexp"
)

type UsersHandler struct {
  writer http.ResponseWriter
  request *http.Request
}

func (handler UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  handler.writer = w
  handler.request = r
  handler.perform()
}

func (handler UsersHandler) perform() {
  parts := strings.Split(handler.request.URL.Path, "/")
  if parts[1] != "users" {
    handler.writer.WriteHeader(http.StatusBadRequest)
    return
  }

  uuidMatcher := regexp.MustCompile(`\A[0-9a-zA-Z]{8}-([0-9a-zA-Z]{4}-){3}[0-9a-zA-Z]{12}\z`)
  if !uuidMatcher.MatchString(parts[2]) {
    handler.writer.WriteHeader(http.StatusBadRequest)
    return
  }

  fmt.Fprintf(handler.writer, "Looking for user '%s'", parts[2])
}
