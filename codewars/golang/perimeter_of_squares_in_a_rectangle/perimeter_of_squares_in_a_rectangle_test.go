package perimeter_of_squares_in_a_rectangle_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(n int, exp int) {
	var ans = Perimeter(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(5, 80)
		dotest(7, 216)
		dotest(20, 114624)
		dotest(30, 14098308)
	})
})
