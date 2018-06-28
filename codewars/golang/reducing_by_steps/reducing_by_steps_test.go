package reducing_by_steps

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(f FParam, arr []int, init int, exp []int) {
	var ans = OperArray(f, arr, init)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases gcdi", func() {
		var dta = []int{18, 69, -90, -78, 65, 40}
		var sol = []int{18, 3, 3, 3, 1, 1}
		dotest(Gcdi, dta, dta[0], sol)
	})
	It("should handle basic cases som", func() {
		var dta = []int{357, 112, 28, -52, 644, 119}
		var sol = []int{357, 469, 497, 445, 1089, 1208}
		dotest(Som, dta, 0, sol)
	})
	It("should handle basic cases maxi", func() {
		var dta = []int{10, -32, 190, 300, -42, -38, 50, 405, -46, 225, -31}
		var sol = []int{10, 10, 190, 300, 300, 300, 300, 405, 405, 405, 405}
		dotest(Maxi, dta, dta[0], sol)
	})
	It("should handle basic cases lcmu", func() {
		var dta = []int{6, -72, -62, -22, -23, 80}
		var sol = []int{6, 72, 2232, 24552, 564696, 5646960}
		dotest(Lcmu, dta, dta[0], sol)
	})
	It("should handle basic cases mini", func() {
		var dta = []int{64, -67, -43, 12, -15, 108, 12, 104, -36}
		var sol = []int{64, -67, -67, -67, -67, -67, -67, -67, -67}
		dotest(Mini, dta, dta[0], sol)
	})
})
