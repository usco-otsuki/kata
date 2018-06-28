package which_x_for_sum

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

func assertFuzzy(m float64, expect float64) {
	var inrange bool
	var merr float64 = 1e-12
	fmt.Printf("m: %0.2f \n", m)
	var actual = Solve(m)
	inrange = (math.Abs(actual-expect) <= merr)
	if inrange == false {
		fmt.Printf("Expected should be near: %0.12e , but got: %0.12e\n", expect, actual)
	}
	Expect(inrange).To(Equal(true))
}

var _ = Describe("Test solve", func() {

	It("should handle basic cases", func() {
		assertFuzzy(2.00, 5.000000000000e-01)
		assertFuzzy(4.00, 6.096117967978e-01)
		assertFuzzy(5.00, 6.417424305044e-01)

	})
})
