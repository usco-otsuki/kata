package integer_partitions

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testequal(n int, exp string) {
	var ans = Part(n)
	Expect(ans).To(Equal(exp))
}

func testEqualEnumPart(n int, exp [][]int) {
	var ans = EnumPart(n)
	Expect(len(ans)).To(Equal(len(exp)))
	for i := 0; i < len(ans); i++ {
		Expect(ans[i]).To(Equal(exp[i]))
	}
}

var _ = Describe("Part", func() {

	It("should handle basic cases John", func() {
		testequal(1, "Range: 0 Average: 1.00 Median: 1.00")
		testequal(2, "Range: 1 Average: 1.50 Median: 1.50")
		testequal(3, "Range: 2 Average: 2.00 Median: 2.00")
		testequal(4, "Range: 3 Average: 2.50 Median: 2.50")
		testequal(5, "Range: 5 Average: 3.50 Median: 3.50")
		testequal(6, "Range: 8 Average: 4.75 Median: 4.50")
	})
})

var _ = Describe("EnumPart", func() {

	It("should handle test cases for EnumPart", func() {
		testEqualEnumPart(1, [][]int{{1}})
		testEqualEnumPart(2, [][]int{{2}, {1, 1}})
		testEqualEnumPart(3, [][]int{{3}, {2, 1}, {1, 1, 1}})
		testEqualEnumPart(4, [][]int{{4}, {3, 1}, {2, 2}, {2, 1, 1}, {1, 1, 1, 1}})
	})
})
