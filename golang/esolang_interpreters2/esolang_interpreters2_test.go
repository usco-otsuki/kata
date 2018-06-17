package esolang_interpreters2

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var examples = []testCase{
	{
		input{"*", "00101100"},
		"10101100",
		"Flips the leftmost cell of the tape",
	},
	{
		input{">*>*", "00101100"},
		"01001100",
		"Flips the second and third cell of the tape",
	},
	{
		input{"*>*>*>*>*>*>*>*", "00101100"},
		"11010011",
		"Flips all the bits in the tape",
	},
	{
		input{"*>*>>*>>>*>*", "00101100"},
		"11111111",
		"Flips all the bits that are initialized to 0",
	},
	{
		input{">>>>>*<*<<*", "00101100"},
		"00000000",
		"Goes somewhere to the right of the tape and then flips all bits that are initialized to 1, progressing leftwards through the tape",
	},
	{
		input{"[>[*>*>*>]>]", "10110"},
		"10101",
		"",
	},
}

var _ = Describe("Smallfuck Interpreter", func() {
	Describe("Examples", func() {
		for i, p := range examples {
			in, out, msg := p.in, p.out, p.msg
			It(fmt.Sprintf("Test #%02d", i+1), func() {
				ans := Interpreter(in.code, in.tape)
				if ans != out {
					fmt.Printf("<LOG::Input>(%#v, %#v)\n", in.code, in.tape)
					fmt.Printf("<LOG::Returned>%#v\n", ans)
					fmt.Printf("<LOG::Expected>%#v\n", out)
					fmt.Println(msg)
				}
				Expect(ans).To(Equal(out))
			})
		}
	})
})

type input struct {
	code string
	tape string
}

type testCase struct {
	in  input
	out string
	msg string
}
