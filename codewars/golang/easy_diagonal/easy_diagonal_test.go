package easy_diagonal_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(n, p int, exp int) {
	var ans = Diagonal(n, p)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(20, 3, 5985)
		dotest(20, 4, 20349)
		dotest(20, 5, 54264)
		dotest(20, 15, 20349)
		dotest(100, 0, 101)
		dotest(100, 10, 158940114100040)
	})
})
