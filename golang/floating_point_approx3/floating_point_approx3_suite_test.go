package floating_point_approx3_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFloatingPointApprox3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FloatingPointApprox3 Suite")
}
