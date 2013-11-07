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

  Describe("finding records", func() {
    It("finds by id", func() {
      db.Exec("INSERT INTO shapes (id, name, sides) VALUES(4, 'pentagon', 5)")
      result := make(map[string]interface{})
      result["id"] = nil
      result["name"] = nil
      result["sides"] = nil
      result, err := cvdb.Find(db, "shapes", 4, result)
      if err != nil {
        fmt.Println(err)
      }
      expectedResult := map[string]interface{}{ "id": int64(4), "name": "pentagon", "sides": int64(5) }
      Expect(result).To(Equal(expectedResult))
    })
  })
})
