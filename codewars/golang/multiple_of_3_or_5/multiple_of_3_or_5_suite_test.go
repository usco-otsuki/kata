package multiple_of_3_or_5_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMultipleOf3Or5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MultipleOf3Or5 Suite")
}
