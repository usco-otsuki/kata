package recreation_two

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(a, b, c, d int, exp [][]int) {
	var ans = Prod2Sum(a, b, c, d)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(1, 2, 1, 3, [][]int{{1, 7}, {5, 5}})
		dotest(2, 3, 4, 5, [][]int{{2, 23}, {7, 22}})
		dotest(1, 2, 2, 3, [][]int{{1, 8}, {4, 7}})
		dotest(1, 1, 3, 5, [][]int{{2, 8}})
		dotest(10, 11, 12, 13, [][]int{{2, 263}, {23, 262}})
		dotest(1, 20, -4, -5, [][]int{{75, 104}, {85, 96}})
	})
})
