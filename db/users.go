package db

type User struct {
  UUID string `db:"uuid"`
}

func (conn Connection) createUsersTable() error {
  createStatement := `CREATE TABLE IF NOT EXISTS users (uuid CHAR(36) NOT NULL PRIMARY KEY);`
  if rows, err := conn.dbConn.Query(createStatement); err != nil {
    return err
  } else {
    rows.Close()
    return nil
  }
}

func (conn Connection) GetUser(uuid string) (*User, bool, error) {
  users := []User{}
  query := "SELECT * FROM users where uuid=?"
  err := conn.dbConn.Select(&users, query, uuid)

  if err != nil {
    return nil, false, err
  } else if len(users) == 0 {
    return nil, false, nil
  } else {
    return &users[0], true, err
  }
}
