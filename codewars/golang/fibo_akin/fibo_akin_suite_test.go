package fibo_akin_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFiboAkin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FiboAkin Suite")
}
