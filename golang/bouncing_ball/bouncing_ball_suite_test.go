package bouncing_ball_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBouncingBall(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BouncingBall Suite")
}
