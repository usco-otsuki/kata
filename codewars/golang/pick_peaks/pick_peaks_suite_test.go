package pick_peaks_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPickPeaks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PickPeaks Suite")
}
