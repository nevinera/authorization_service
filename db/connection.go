package db

import (
  "github.com/jmoiron/sqlx"
  _ "github.com/go-sql-driver/mysql"
)

type Connection struct {
  configString string
  dbConn *sqlx.DB
}

func NewConnection(configString string) (*Connection, error) {
  conn := Connection{configString: configString}
  dbConn := sqlx.MustConnect("mysql", configString)
  conn.dbConn = dbConn
  return &Connection{configString: configString, dbConn: dbConn}, nil
}

func (conn Connection) CreateDatabase() {
  conn.createUsersTable()
}
