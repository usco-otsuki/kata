package fibo_akin

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest1(n, k int, exp int) {
	var ans = LengthSupUk(n, k)
	Expect(ans).To(Equal(exp))
}
func dotest2(n int, exp int) {
	var ans = Comp(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases LengthSupUk", func() {
		dotest1(50, 25, 2)
		dotest1(3332, 973, 1391)
	})
	It("should handle basic cases Comp", func() {
		dotest2(74626, 37128)
		dotest2(71749, 35692)
	})
})
