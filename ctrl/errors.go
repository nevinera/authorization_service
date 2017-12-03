package ctrl

import (
  "net/http"
  "strconv"
  "time"
  "encoding/json"
)

func sendError(w http.ResponseWriter, status int, msg string) {
  w.WriteHeader(status)
  errorData := map[string]string{
    "error": msg,
    "status": strconv.Itoa(status),
    "occurred_at": time.Now().String(),
  }
  json.NewEncoder(w).Encode(errorData)
}
