package basee91

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// EncodeToString encodes the given byte array and returns a string
func EncodeToString(d []byte) string {
	return string(Encode(d))
}

// DecodeToString decodes a given byte array are returns a string
func DecodeToString(d []byte) string {
	return string(Decode(d))
}

// var _ = Describe("Decoding", func() {
// 	It("should pass example tests", func() {
// 		Expect(DecodeToString([]byte("fPNKd"))).To(Equal("test"))
// 		Expect(DecodeToString([]byte(">OwJh>Io0Tv!8PE"))).To(Equal("Hello World!"))
// 	})
// })

var _ = Describe("Encoding", func() {
	It("should pass example tests", func() {
		Expect(EncodeToString([]byte("A"))).To(Equal("A%"))
		Expect(EncodeToString([]byte("test"))).To(Equal("fPNKd"))
		Expect(EncodeToString([]byte("Hello World!"))).To(Equal(">OwJh>Io0Tv!8PE"))
	})
})
