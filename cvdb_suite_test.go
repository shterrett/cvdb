package cvdb_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCvdb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cvdb Suite")
}
