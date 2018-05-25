package convert_string_to_camel_case

import (
	. "fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(str, exp string) {
	Println("input:", str)
	var ans = ToCamelCase(str)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {
	It("should handle basic cases", func() {
		dotest("", "")
		dotest("The_Stealth_Warrior", "TheStealthWarrior")
		dotest("the-Stealth-Warrior", "theStealthWarrior")
	})
})
