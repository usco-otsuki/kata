package financing_plan_on_planet

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testequal(n int, exp int) {
	var ans = Finance(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		testequal(5, 105)
		testequal(6, 168)
		testequal(8, 360)
		testequal(15, 2040)
	})
})
