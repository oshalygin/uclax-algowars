package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("PlusMinus", func() {
	It("Input: [[-4, 3, -9, 0, 4, 1]] => 0.500000 0.333333 0.166667", func() {

		input := []int{-4, 3, -9, 0, 4, 1}

		expected := PlusMinusFractions{Positive: 0.500000, Negative: 0.3333333333333333, Zero: 0.16666666666666666}

		actual := PlusMinus(input)
		Expect(actual).To(Equal(expected))
	})

	It("Input: [[4, 3, 9, 2, 4, 1]] => 1 0 0", func() {

		input := []int{4, 3, 9, 2, 4, 1}

		expected := PlusMinusFractions{Positive: 1, Negative: 0, Zero: 0}

		actual := PlusMinus(input)
		Expect(actual).To(Equal(expected))
	})
})
