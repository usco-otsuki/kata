package disease_spread

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testequal(tm, n, s0, i0 int, b, a float64, exp int) {
	var ans = Epidemic(tm, n, s0, i0, b, a)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		testequal(18, 432, 1004, 1, 0.00209, 0.51, 420)
		testequal(12, 288, 1007, 2, 0.00206, 0.45, 461)
		testequal(13, 312, 999, 1, 0.00221, 0.55, 409)
	})
})
