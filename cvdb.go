package cvdb

import (
  _ "github.com/lib/pq"
  "database/sql"
  "fmt"
  "strings"
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
    result[v] = Cast(values[i])
  }
  return result, nil
}

func FindAll(db *sql.DB, table string, result []map[string]interface{}) ([]map[string]interface{}, error) {
  queryString := fmt.Sprintf("SELECT * FROM %s", table)
  rows, err := db.Query(queryString)
  if err != nil {
    return nil, err
  }
  columns, err := rows.Columns()
  if err != nil {
    return nil, err
  }
  for rows.Next() == true {
    record := make([]interface{}, len(columns))
    recordPointer := make([]interface{}, len(columns))
    for i, _ := range columns {
      record[i] = nil
      recordPointer[i] = &record[i]
    }
    rows.Scan(recordPointer...)
    rowMap := make(map[string]interface{})
    for i, v := range columns {
      rowMap[v] = Cast(record[i])
    }
    result = append(result, rowMap)
  }
  return result, nil
}
