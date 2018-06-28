package lcs

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testLCS(s1, s2, s string) {
	var ans = LCS(s1, s2)
	Expect(ans).To(Equal(s))
}

var _ = Describe("Sample Test Cases", func() {

	It("LCS should work on sample case #1", func() {
		testLCS("a", "b", "")
	})

	It("LCS should work on sample case #2", func() {
		testLCS("abcdef", "abc", "abc")
	})

	It("LCS should work on sample case #3", func() {
		testLCS("132535365", "123456789", "12356")
	})
})
