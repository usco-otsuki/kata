package disease_spread_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDiseaseSpread(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DiseaseSpread Suite")
}
