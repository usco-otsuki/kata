package split_strings_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSplitStrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SplitStrings Suite")
}
