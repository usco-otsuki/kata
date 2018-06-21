package lcs_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLcs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lcs Suite")
}
