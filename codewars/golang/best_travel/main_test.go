package best_travel

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(t, k int, ls []int, exp int) {
	var ans = ChooseBestSum(t, k, ls)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		var ts = []int{50, 55, 56, 57, 58}
		dotest(163, 3, ts, 163)
		ts = []int{50}
		dotest(163, 3, ts, -1)
		ts = []int{91, 74, 73, 85, 73, 81, 87}
		dotest(230, 3, ts, 228)
		dotest(331, 2, ts, 178)
		dotest(331, 4, ts, 331)
		dotest(331, 5, ts, -1)
	})
})
