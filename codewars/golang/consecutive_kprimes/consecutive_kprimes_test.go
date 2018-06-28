package consecutive_kprimes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(k int, arr []int, exp int) {
	var ans = ConsecKprimes(k, arr)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("isPrevKPrime", func() {

	It("should handle basic cases", func() {
		Expect(isKPrime(10, 2)).To(Equal(true))
		Expect(isKPrime(16, 4)).To(Equal(true))
		Expect(isKPrime(17, 2)).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		// dotest(2, []int{10081, 10071, 10077, 10065, 10060, 10070, 10086, 10083, 10078, 10076, 10089, 10085, 10063, 10074, 10068, 10073, 10072, 10075}, 2)
		dotest(6, []int{10064}, 0)
		dotest(1, []int{10054, 10039, 10053, 10051, 10047, 10043, 10037, 10034}, 0)
		dotest(3, []int{10158, 10182, 10184, 10172, 10179, 10168, 10156, 10165, 10155, 10161, 10178, 10170}, 5)
		dotest(2, []int{10110, 10102, 10097, 10113, 10123, 10109, 10118, 10119, 10117, 10115, 10101, 10121, 10122}, 7)
	})
})
