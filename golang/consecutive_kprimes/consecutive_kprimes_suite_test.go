package consecutive_kprimes_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConsecutiveKprimes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConsecutiveKprimes Suite")
}
