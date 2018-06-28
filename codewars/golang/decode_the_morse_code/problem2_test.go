package decode_the_morse_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EncodeMorse", func() {
	It("HEY JUDE", func() {
		Expect(EncodeMorse("HEY JUDE")).To(Equal(".... . -.--   .--- ..- -.. ."))
	})
})

var _ = Describe("FindMaxOneLen", func() {
	It("000011100001111111000", func() {
		Expect(FindMaxOneLen("000011100001111111000")).To(Equal(7))
	})
})

var _ = Describe("DecodeBits", func() {
	It("outputs empty string for 000", func() {
		Expect(DecodeBits("000")).To(Equal(""))
	})

	It("prioritize dot on dash when there is an ambiguity", func() {
		Expect(DecodeBits("111")).To(Equal("."))
	})

	It("trims preceding and trailing 0s", func() {
		Expect(DecodeBits("001100")).To(Equal("."))
	})

	It("", func() {
		Expect(DecodeBits("111000111111111")).To(Equal(".-"))
	})

	It("", func() {
		Expect(DecodeBits("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011")).To(Equal(".... . -.--   .--- ..- -.. ."))
	})

})

var _ = Describe("DecodeMorse", func() {
	It("Example from description", func() {
		Expect(DecodeMorse(DecodeBits("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"))).To(Equal("HEY JUDE"))
	})

	It("A", func() {
		Expect(DecodeMorse(DecodeBits("111000111111111"))).To(Equal("A"))
	})
})
