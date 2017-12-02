package main

import (
  "fmt"
  "net/http"
  "./handlers"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Fake HTTP response to `%s`", r.URL.Path)
}

func main() {
  http.HandleFunc("/", handler)
  http.Handle("/users/", handlers.UsersHandler{})
  http.ListenAndServe(":3000", nil)
}
