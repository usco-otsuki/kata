package twice_linear

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(n int, exp int) {
	var ans = DblLinear(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(50, 175)
		dotest(100, 447)
		dotest(500, 3355)
		dotest(1000, 8488)
	})
})
