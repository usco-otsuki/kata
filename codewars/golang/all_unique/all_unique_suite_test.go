package all_unique_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAllUnique(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AllUnique Suite")
}
