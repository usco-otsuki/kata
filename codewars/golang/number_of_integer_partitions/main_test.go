package number_of_integer_partitions

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample Test Cases", func() {
	It("the number of integer partitions of 1 should be correct", func() {
		Expect(Partitions(1)).To(Equal(1))
	})

	It("the number of integer partitions of 5 should be correct", func() {
		Expect(Partitions(5)).To(Equal(7))
	})

	It("the number of integer partitions of 10 should be correct", func() {
		Expect(Partitions(10)).To(Equal(42))
	})

	It("the number of integer partitions of 25 should be correct", func() {
		Expect(Partitions(25)).To(Equal(1958))
	})
})
