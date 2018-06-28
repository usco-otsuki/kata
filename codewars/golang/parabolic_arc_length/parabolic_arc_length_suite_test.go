package parabolic_arc_length_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestParabolicArcLength(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ParabolicArcLength Suite")
}
