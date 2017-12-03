package db

import (
  "github.com/jmoiron/sqlx"
  _ "github.com/go-sql-driver/mysql"
)

type Connection struct {
  configString string
  dbConn *sqlx.DB
  connected bool
}

func NewConnection(configString string) (Connection, error) {
  conn := Connection{configString: configString}
  if dbConn, err := sqlx.Connect("mysql", configString); err != nil {
    conn.connected = false
    conn.dbConn = nil
    return conn, err
  } else {
    conn.connected = true
    conn.dbConn = dbConn
    return Connection{configString: configString, dbConn: dbConn}, nil
  }
}

func (conn Connection) CreateDatabase() {
  conn.createUsersTable()
}
