package diophantine_equation_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDiophantineQuation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DiophantineQuation Suite")
}
