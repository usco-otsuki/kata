package binary_multiple_of_3

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"regexp"
	"strconv"
)

var _ = Describe("Test regexp", func() {

	type test struct {
		value  string
		result bool
	}
	simpleTests := []test{{" 0", false}, {"abc", false}, {"000", true}, {"110", true}, {"111", false}, {strconv.FormatInt(12345678, 2), true}, {"123", false}, {" 000  ", false}}
	for _, t := range simpleTests {

		It("should test that the solution returns the correct value for "+t.value, func() {
			matched, _ := regexp.Match(MultipleOf3Regex, []byte(t.value))
			Expect(matched).To(Equal(t.result))
		})
	}
})
