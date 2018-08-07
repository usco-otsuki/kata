package reversed_string

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("simple word", func() {
		Expect(Solution("world")).To(Equal("dlrow"))
	})
	It("Empty string", func() {
		Expect(Solution("")).To(Equal(""))
	})
	It("one character", func() {
		Expect(Solution("a")).To(Equal("a"))
	})
})
