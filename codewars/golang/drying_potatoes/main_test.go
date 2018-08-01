package drying_potatoes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(p0, w0, p1 int, exp int) {
	var ans = Potatoes(p0, w0, p1)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(99, 100, 98, 50)
		dotest(82, 127, 80, 114)
	})
})
