package drying_potatoes_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDryingPotatoes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DryingPotatoes Suite")
}
