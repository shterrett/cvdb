package cvdb_test

import (
	"github.com/shterrett/cvdb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
  "database/sql"
  "fmt"
)

var _ = Describe("cvdbInitialize", func() {
  It("connects to a database", func() {
    db, err := cvdb.ConnectTo("stuart", "gothings", "disable")
    Expect(err).To(BeNil())
    err = db.Ping()
    Expect(err).To(BeNil())
  })
})

var _ = Describe("helpers", func() {
  It("generates a slice of column names", func() {
    dbRecord := map[string]interface{}{ "oneColumn": "oneValue", "twoColumn": "twoValue" }
    columnNames := cvdb.ColNames(dbRecord)
    expectedColumnNames := []string{ "oneColumn", "twoColumn" }
    Expect(columnNames).To(Equal(expectedColumnNames))
  })

  It("generates a slice of column values in same orer as names", func() {
    dbRecord := map[string]interface{}{ "oneColumn": "oneValue", "twoColumn": "twoValue" }
    columnNames := cvdb.ColNames(dbRecord)
    columnArgs := cvdb.ColArgs(dbRecord, columnNames)
    expectedColumnArgs := []interface{}{ "oneValue", "twoValue" }
    Expect(columnArgs).To(Equal(expectedColumnArgs))
  })

  It("generates a string with a placeholder for each column", func() {
    dbRecord := map[string]interface{}{ "oneColumn": "oneValue", "twoColumn": "twoValue" }
    columnNames := cvdb.ColNames(dbRecord)
    placeholders := cvdb.Placeholders(columnNames)
    expectedPlaceholders := "$1, $2"
    Expect(placeholders).To(Equal(expectedPlaceholders))
  })
})

var _ = Describe("database operations", func() {
  var db *sql.DB
  newShape := make(map[string]interface{})

  BeforeEach(func() {
    var err error
    db, err = cvdb.ConnectTo("stuart", "gothings", "disable")
    if err != nil {
      panic("No db, no testy")
    }

    newShape["name"] = "square"
    newShape["sides"] = 4

    db.Exec("TRUNCATE shapes")
  })

  Describe("creating records", func() {
    It("inserts a record into the database", func() {
      err := cvdb.Create(db, "shapes", newShape)
      if err != nil {
        fmt.Println(err)
      }
      queryResults, _ := db.Query("SELECT *  FROM shapes WHERE name = 'square'")
      var id int
      var sides int
      var name string
      queryResults.Next()
      queryResults.Scan(&id, &sides, &name)
      Expect(name).To(Equal(newShape["name"]))
      Expect(sides).To(Equal(newShape["sides"]))
      queryResults.Close()
    })
  })
})
