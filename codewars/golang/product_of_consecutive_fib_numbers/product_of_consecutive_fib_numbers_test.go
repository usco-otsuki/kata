package product_of_consecutive_fib_numbers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(prod uint64, exp [3]uint64) {
	var ans = ProductFib(prod)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(4895, [3]uint64{55, 89, 1})
		dotest(5895, [3]uint64{89, 144, 0})
		dotest(74049690, [3]uint64{6765, 10946, 1})
		dotest(84049690, [3]uint64{10946, 17711, 0})
	})
})
