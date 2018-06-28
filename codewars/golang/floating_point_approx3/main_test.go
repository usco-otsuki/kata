package floating_point_approx3

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

func assertFuzzyEquals(a float64, b float64, c float64) {
	var inrange bool
	var merr float64 = 1e-12
	fmt.Printf("a: %1.6e b: %1.6e c: %1.6e\n", a, b, c)
	var x = Quadratic(a, b, c)
	var smallest = math.Abs(x) < 1.e-1
	if smallest == false {
		fmt.Printf("This root is not the good one: %1.6e \n", x)
	}
	var y = a*x*x + b*x + c
	inrange = (math.Abs(y) <= merr)
	if inrange == false {
		fmt.Printf("Expected f(x) should be near: 0 , but got: %1.6e\n", y)
	}
	//fmt.Printf("%t \n", inrange && smallest)
	Expect(inrange && smallest).To(Equal(true))
}

var _ = Describe("Test quadratic", func() {

	It("should handle basic cases", func() {
		assertFuzzyEquals(7, 4.00e+13, 8)
		assertFuzzyEquals(9, 1.00e+14, 1)
		assertFuzzyEquals(3, 3.00e+09, 1)
		assertFuzzyEquals(7, 4.00e+09, 7)

	})

})
