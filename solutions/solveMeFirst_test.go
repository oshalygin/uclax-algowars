package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("SolveMeFirst", func() {
	It("Input: 2 + 2 => 4", func() {
		x := 2
		y := 2

		expected := 4
		actual := SolveMeFirst(x, y)

		Expect(actual).To(Equal(expected))
	})
})
