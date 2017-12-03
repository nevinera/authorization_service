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

func (conn Connection) getUser(uuid string) (User, error) {
  user := User{}
  query := "SELECT * FROM users where uuid = $1"

  if err := conn.dbConn.Get(&user, query, uuid); err != nil {
    return user, err
  } else {
    return user, nil
  }
}
