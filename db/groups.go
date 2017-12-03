package db

type Group struct {
  UUID string `db:"uuid"`
}

func (conn Connection) createGroupsTable() error {
  createStatement := `CREATE TABLE IF NOT EXISTS groups (uuid CHAR(36) NOT NULL PRIMARY KEY);`
  if rows, err := conn.dbConn.Query(createStatement); err != nil {
    return err
  } else {
    rows.Close()
    return nil
  }
}

func (conn Connection) GetGroup(uuid string) (*Group, bool, error) {
  groups := []Group{}
  query := "SELECT * FROM groups where uuid=?"
  err := conn.dbConn.Select(&groups, query, uuid)

  if err != nil {
    return nil, false, err
  } else if len(groups) == 0 {
    return nil, false, nil
  } else {
    return &groups[0], true, err
  }
}

func (conn Connection) CreateGroup(uuid string) (*Group, bool, error) {
  group, found, fetchErr := conn.GetGroup(uuid)
  if fetchErr != nil {
    return group, false, fetchErr
  } else if found == true {
    return group, false, nil
  }

  createdGroup := &Group{UUID: uuid}
  stmt := "INSERT INTO groups (uuid) VALUES (:uuid)"
  varMap := map[string]interface{}{"uuid": uuid}
  _, createErr := conn.dbConn.NamedQuery(stmt, varMap)
  if createErr != nil {
    return createdGroup, false, createErr
  } else {
    return createdGroup, true, nil
  }
}

func (conn Connection) DestroyGroup(uuid string) (*Group, bool, error) {
  group, found, fetchErr := conn.GetGroup(uuid)
  if fetchErr != nil {
    return nil, false, fetchErr
  } else if found == false {
    return nil, false, nil
  }

  stmt := "DELETE FROM groups WHERE uuid = ?"
  _, deleteErr := conn.dbConn.Query(stmt, uuid)
  if deleteErr != nil {
    return group, false, deleteErr
  } else {
    return group, true, nil
  }
}
