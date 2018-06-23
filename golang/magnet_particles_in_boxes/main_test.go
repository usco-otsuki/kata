package magnet_particles_in_boxes

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

func assertFuzzyEquals(act float64, exp float64) {
	var inrange bool
	var merr float64 = 1e-12
	var e float64
	if exp == 0.0 {
		e = math.Abs(act)
	} else {
		e = math.Abs((act - exp) / exp)
	}
	inrange = (e <= merr)
	if inrange == false {
		fmt.Printf("Expected should be near: %1.12e , but got: %1.12e\n", exp, act)
	}
	Expect(inrange).To(Equal(true))
}
func dotest(maxk, maxn int, exp float64) {
	assertFuzzyEquals(Doubles(maxk, maxn), exp)
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(1, 10, 0.5580321939764581)
		dotest(10, 1000, 0.6921486500921933)
		dotest(10, 10000, 0.6930471674194457)
	})
})
