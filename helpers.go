package cvdb

import (
  _ "github.com/lib/pq"
  "fmt"
  "strings"
  "reflect"
)

func Cast(value interface{}) interface{} {
  var result interface{}
  if reflect.TypeOf(value).String() == "[]uint8" {
    result = string(value.([]uint8))
  } else if reflect.TypeOf(value).String() == "int64" {
    result = int64(value.(int64))
  }
  return result
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
