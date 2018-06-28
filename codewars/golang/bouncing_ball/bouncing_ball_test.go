package bouncing_ball

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testequal(h, bounce, window float64, exp int) {
	var ans = BouncingBall(h, bounce, window)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		testequal(3, 0.66, 1.5, 3)
		testequal(40, 0.4, 10, 3)
		testequal(10, 0.6, 10, -1)
		testequal(40, 1, 10, -1)
		testequal(5, -1, 1.5, -1)
	})
})
