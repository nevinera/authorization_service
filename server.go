package main

import (
  "os"
  "log"
  "github.com/nevinera/authorization_service/server"
)

func main() {
  dbUrl, dbWasSet := os.LookupEnv("DATABASE_URL")
  if !dbWasSet {
    dbUrl = "root:@tcp(localhost:3306)/authorization"
  }

  listenString, listenWasSet := os.LookupEnv("LISTEN")
  if !listenWasSet {
    listenString = ":3000"
  }

  err := server.Run(listenString, dbUrl)
  if err != nil {
    log.Fatal(err)
  }
}
