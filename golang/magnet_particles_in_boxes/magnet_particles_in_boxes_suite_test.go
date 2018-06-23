package magnet_particles_in_boxes_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMagnetParticlesInBoxes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MagnetParticlesInBoxes Suite")
}
