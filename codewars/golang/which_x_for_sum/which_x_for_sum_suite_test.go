package which_x_for_sum_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWhichXForSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WhichXForSum Suite")
}
