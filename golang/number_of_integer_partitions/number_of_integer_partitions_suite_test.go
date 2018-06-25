package number_of_integer_partitions_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNumberOfIntegerPartitions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NumberOfIntegerPartitions Suite")
}
