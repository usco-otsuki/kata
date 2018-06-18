package best_travel_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBestTravel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BestTravel Suite")
}
