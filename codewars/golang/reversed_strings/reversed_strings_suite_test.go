package reversed_string_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReversedStrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ReversedStrings Suite")
}
