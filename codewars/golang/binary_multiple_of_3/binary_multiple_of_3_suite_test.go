package binary_multiple_of_3_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBinaryMultipleOf3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BinaryMultipleOf3 Suite")
}
