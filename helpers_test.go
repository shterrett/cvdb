package cvdb_test

import (
	"github.com/shterrett/cvdb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
  "time"
)

var _ = Describe("helpers", func() {

  dbRecord := map[string]interface{}{ "oneColumn": "oneValue", "twoColumn": "twoValue" }

  It("generates a slice of column names", func() {
    columnNames := cvdb.ColNames(dbRecord)
    expectedColumnNames := []string{ "oneColumn", "twoColumn" }
    Expect(columnNames).To(Equal(expectedColumnNames))
  })

  It("generates a slice of column values in same orer as names", func() {
    columnNames := cvdb.ColNames(dbRecord)
    columnArgs := cvdb.ColArgs(dbRecord, columnNames)
    expectedColumnArgs := []interface{}{ "oneValue", "twoValue" }
    Expect(columnArgs).To(Equal(expectedColumnArgs))
  })

  It("generates a string with a placeholder for each column", func() {
    columnNames := cvdb.ColNames(dbRecord)
    placeholders := cvdb.Placeholders(columnNames)
    expectedPlaceholders := "$1, $2"
    Expect(placeholders).To(Equal(expectedPlaceholders))
  })
})

var _ = Describe("Type inference", func() {
  testValues := make(map[string]interface{})
  testValues["int64"] = int64(5)
  testValues["float64"] = float64(3.14)
  testValues["bool"] = bool(true)
  testValues["[]byte"] = []byte("test byte string")
  testValues["string"] = string("test string")
  testValues["time.Time"] = time.Now()
  testValues["nil"] = new(interface{})

  It("returns an int64 as an int64", func() {
    var result interface{}
    result = cvdb.Cast(interface{}(testValues["int64"]))
    Expect(result).To(Equal(testValues["int64"]))
  })

  It("returns a []byte as a string", func() {
    var result interface{}
    result = cvdb.Cast(interface{}(testValues["[]byte"]))
    Expect(result).To(Equal("test byte string"))
  })
})
