package product_of_consecutive_fib_numbers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProductOfConsecutiveFibNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductOfConsecutiveFibNumbers Suite")
}
