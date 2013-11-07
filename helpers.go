package cvdb

import (
  _ "github.com/lib/pq"
  "fmt"
  "strings"
)

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
