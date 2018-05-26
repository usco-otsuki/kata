package convert_string_to_camel_case_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConvertStringToCamelCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConvertStringToCamelCase Suite")
}
