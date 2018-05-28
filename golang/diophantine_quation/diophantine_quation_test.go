package diophantine_quation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(n int, exp [][]int) {
	var ans = Solequa(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(5, [][]int{{3, 1}})
		dotest(12, [][]int{{4, 1}})
		dotest(13, [][]int{{7, 3}})
		dotest(9005, [][]int{{4503, 2251}, {903, 449}})
		dotest(9008, [][]int{{1128, 562}})
		dotest(90002, [][]int{})
	})
})
