package gap_in_primes_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGapInPrimes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GapInPrimes Suite")
}
