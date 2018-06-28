package range_extraction

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Range Extraction", func() {
	It("should match the example", func() {
		s := Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20})
		Expect(s).To(Equal("-6,-3-1,3-5,7-11,14,15,17-20"))
	})
})
