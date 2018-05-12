package parabolic_arc_length

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

func assertFuzzyEquals(act float64, exp float64) {
	var inrange bool
	var merr float64 = 1e-6
	var e float64
	if exp == 0.0 {
		e = math.Abs(act)
	} else {
		e = math.Abs((act - exp) / exp)
	}
	inrange = (e <= merr)
	if inrange == false {
		fmt.Printf("Expected should be near: %1.6e , but got: %1.6e\n", exp, act)
	}
	Expect(inrange).To(Equal(true))
}
func dotest(n int, exp float64) {
	assertFuzzyEquals(LenCurve(n), exp)
}

var _ = Describe("Test Example", func() {
	It("should handle basic cases", func() {
		dotest(1, 1.414213562)
		dotest(10, 1.478197397)
		dotest(40, 1.478896272)
		dotest(200, 1.478940994)
	})
})
