package decode_the_morse_code_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDecodeTheMorseCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DecodeTheMorseCode Suite")
}
