package cvdb

import (
  _ "github.com/lib/pq"
  "database/sql"
  "fmt"
  "strings"
  "reflect"
)

func ConnectTo(username string, database string, sslmode string) (dbConn *sql.DB, err error) {
  connectionString := fmt.Sprintf("user='%s' dbname='%s' sslmode='%s'", username, database, sslmode)
  dbConn, err = sql.Open("postgres", connectionString)
  return dbConn, err
}

func Create(db *sql.DB, table string, record map[string]interface{}) (err error) {
  columns := ColNames(record)
  columnArgs := ColArgs(record, columns)
  placeholders := Placeholders(columns)
  queryString := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", table, strings.Join(columns, ", "), placeholders)
  _, err = db.Exec(queryString, columnArgs...)
  return err
}

func Find(db *sql.DB, table string, id int, result map[string]interface{}) (map[string]interface{}, error) {
  queryString := fmt.Sprintf("SELECT * FROM %s WHERE id = %d", table, id)
  row, err := db.Query(queryString)
  if err != nil {
    return nil, err
  }
  columns, err := row.Columns()
  if err != nil {
    return nil, err
  }
  values := make([]interface{}, len(columns))
  valuePtrs := make([]interface{}, len(columns))
  for i, _ := range columns {
    valuePtrs[i] = &values[i]
  }
  row.Next()
  row.Scan(valuePtrs...)
  for i, v := range columns {
    if reflect.TypeOf(values[i]).String() == "[]uint8" {
      result[v] = string(values[i].([]uint8))
    } else if reflect.TypeOf(values[i]).String() == "int64" {
      result[v] = int64(values[i].(int64))
    }
  }
  return result, nil
}
