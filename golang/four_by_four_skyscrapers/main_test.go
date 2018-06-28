package four_by_four_skyscrapers

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var clues = []int{
	0, 0, 1, 2,
	0, 2, 0, 0,
	0, 3, 0, 0,
	0, 1, 0, 0}

var outcome = [][]int{
	{2, 1, 4, 3},
	{3, 4, 1, 2},
	{4, 2, 3, 1},
	{1, 3, 2, 4}}

func dotest(clues []int, exp [][]int) {
	fmt.Println("Input", clues)
	Expect(SolvePuzzle(clues)).To(Equal(exp))
}

var _ = Describe("Sample test case", func() {
	It("should solve static puzzle", func() {
		dotest(clues, outcome)
	})
})

var _ = Describe("Sample test case", func() {
	It("should solve static puzzle", func() {
		Expect(IsValid(outcome, clues)).To(Equal(true))
	})
})
