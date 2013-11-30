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

func Find(db *sql.DB, table string, id int) (map[string]interface{}, error) {
  queryString := fmt.Sprintf("SELECT * FROM %s WHERE id = %d", table, id)
  row, err := db.Query(queryString)
  if err != nil {
    return nil, err
  }
  columns, err := row.Columns()
  if err != nil {
    return nil, err
  }
  row.Next()
  record := makeRecord(columns, row)
  return record, nil
}

func FindAll(db *sql.DB, table string) ([]map[string]interface{}, error) {
  queryString := fmt.Sprintf("SELECT * FROM %s", table)
  rows, err := db.Query(queryString)
  if err != nil {
    return nil, err
  }
  columns, err := rows.Columns()
  if err != nil {
    return nil, err
  }
  result := make([]map[string]interface{}, 0)
  for rows.Next() == true {
    record := makeRecord(columns, rows)
    result = append(result, record)
  }
  return result, nil
}

func FindAllWhere(db *sql.DB, table string, column string, value string) ([]map[string]interface{}, error) {
  queryString := fmt.Sprintf("SELECT * FROM %s WHERE %s = %s", table, column, value)
  fmt.Println(queryString)
  rows, err := db.Query(queryString)
  if err != nil {
    return nil, err
  }
  columns, err := rows.Columns()
  if err != nil {
    return nil, err
  }
  result := make([]map[string]interface{}, 0)
  for rows.Next() == true {
    record := makeRecord(columns, rows)
    result = append(result, record)
  }
  return result, nil
}

func makeRecord(columns []string, row *sql.Rows) map[string]interface{} {
  record := make([]interface{}, len(columns))
  recordPointer := make([]interface{}, len(columns))
  for i, _ := range columns {
    record[i] = nil
    recordPointer[i] = &record[i]
  }
  row.Scan(recordPointer...)
  rowMap := make(map[string]interface{})
  for i, v := range columns {
    rowMap[v] = Cast(record[i])
  }
  return rowMap
}
