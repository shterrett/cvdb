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

func ColNames(record map[string]interface{}) []string {
  columnNames := make([]string, 0)
  for k, _ := range record {
    columnNames = append(columnNames, k)
  }
  return columnNames
}

func ColArgs(record map[string]interface{}, columns []string) []interface{} {
  columnArgs := make([]interface{}, 0)
  for _, v := range columns {
    columnArgs = append(columnArgs, record[v])
  }
  return columnArgs
}

func Placeholders(columns []string) string {
  var placeholder string
  var placeholders []string
  for i, _ := range columns {
    placeNumber := i +  1
    placeholder = fmt.Sprintf("$%d", placeNumber)
    placeholders = append(placeholders, placeholder)
  }
  return strings.Join(placeholders, ", ")
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
      result[v] = int(values[i].(int64))
    }
  }
  return result, nil
}

// func insertShape(sides int, name string) (query string) {
//  return fmt.Sprintf("INSERT INTO shapes (sides, name) VALUES('%d', '%s')",
//   sides, name)
// }
//
// func selectShape(name string) (query string) {
//   return fmt.Sprintf("SELECT * FROM shapes WHERE name = '%s'",
//     name)
// }
//
// func main() {
//   dbPg, err := sql.Open("postgres", "user=stuart dbname=gothings sslmode=disable")
//   if err != nil {
//     fmt.Println(err)
//   }
//   _, err = dbPg.Exec(insertShape(4, "square"))
//   if err != nil {
//     fmt.Println(err)
//   }
//   queryResults, err := dbPg.Query(selectShape("square"))
//   if err != nil {
//     fmt.Println(err)
//   }
//   for queryResults.Next() {
//     fmt.Println(shape_id, sides, name)
//   }
// }
