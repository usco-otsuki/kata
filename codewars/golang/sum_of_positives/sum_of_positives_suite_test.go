package sum_of_positives_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSumOfPositives(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SumOfPositives Suite")
}
