package john_and_ann_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJohnAndAnn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JohnAndAnn Suite")
}
