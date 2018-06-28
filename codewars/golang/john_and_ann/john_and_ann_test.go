package john_and_ann

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotestJohn(n int, exp []int) {
	var ans = John(n)
	Expect(ans).To(Equal(exp))
}
func dotestAnn(n int, exp []int) {
	var ans = Ann(n)
	Expect(ans).To(Equal(exp))
}
func dotestSumJohn(n int, exp int) {
	var ans = SumJohn(n)
	Expect(ans).To(Equal(exp))
}
func dotestSumAnn(n int, exp int) {
	var ans = SumAnn(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases John", func() {
		dotestJohn(11, []int{0, 0, 1, 2, 2, 3, 4, 4, 5, 6, 6})
	})

	It("should handle basic cases Ann", func() {
		dotestAnn(6, []int{1, 1, 2, 2, 3, 3})
	})
	It("should handle basic cases SumJohn", func() {
		dotestSumJohn(75, 1720)
	})
	It("should handle basic cases SumAnn", func() {
		dotestSumAnn(115, 4070)
	})
})
