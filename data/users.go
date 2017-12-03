package data

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

func (conn Connection) CreateUser(uuid string) (*User, bool, error) {
  user, found, fetchErr := conn.GetUser(uuid)
  if fetchErr != nil {
    return user, false, fetchErr
  } else if found == true {
    return user, false, nil
  }

  createdUser := &User{UUID: uuid}
  stmt := "INSERT INTO users (uuid) VALUES (:uuid)"
  varMap := map[string]interface{}{"uuid": uuid}
  _, createErr := conn.dbConn.NamedQuery(stmt, varMap)
  if createErr != nil {
    return createdUser, false, createErr
  } else {
    return createdUser, true, nil
  }
}

func (conn Connection) DestroyUser(uuid string) (*User, bool, error) {
  user, found, fetchErr := conn.GetUser(uuid)
  if fetchErr != nil {
    return nil, false, fetchErr
  } else if found == false {
    return nil, false, nil
  }

  stmt := "DELETE FROM users WHERE uuid = ?"
  _, deleteErr := conn.dbConn.Query(stmt, uuid)
  if deleteErr != nil {
    return user, false, deleteErr
  } else {
    return user, true, nil
  }
}
