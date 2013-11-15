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
      result, err := cvdb.Find(db, "shapes", 4)
      if err != nil {
        fmt.Println(err)
      }
      expectedResult := map[string]interface{}{ "id": int64(4), "name": "pentagon", "sides": int64(5) }
      Expect(result).To(Equal(expectedResult))
    })

    It("finds all records in a table", func() {
      db.Exec("INSERT INTO shapes (id, name, sides) VALUES(1, 'pentagon', 5)")
      db.Exec("INSERT INTO shapes (id, name, sides) VALUES(2, 'square', 4)")
      result, err := cvdb.FindAll(db, "shapes")
      if err != nil {
        fmt.Println(err)
      }
      shape1 := make(map[string]interface{})
      shape1["id"] = int64(1)
      shape1["sides"] = int64(5)
      shape1["name"] = "pentagon"
      shape2 := make(map[string]interface{})
      shape2["id"] = int64(2)
      shape2["sides"] = int64(4)
      shape2["name"] = "square"
      expectedResult := []map[string]interface{}{shape1, shape2}
      Expect(result).To(Equal(expectedResult))
    })
  })
})
