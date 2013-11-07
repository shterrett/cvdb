package cvdb_test

import (
	"github.com/shterrett/cvdb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
