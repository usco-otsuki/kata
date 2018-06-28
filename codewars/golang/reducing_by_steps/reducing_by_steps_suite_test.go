package reducing_by_steps_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReducingBySteps(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ReducingBySteps Suite")
}
