package solutions_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSolutions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Solutions Suite")
}
