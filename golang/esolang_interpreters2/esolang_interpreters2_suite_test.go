package esolang_interpreters2_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEsolangInterpreters2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EsolangInterpreters2 Suite")
}
