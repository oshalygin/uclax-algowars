package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/oshalygin/uclax-algowars/solutions"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("MiniMaxSum", func() {
	It("1 2 3 4 5 => 10 14", func() {
		input := []int64{1, 2, 3, 4, 5}
		expected := solutions.MiniMax{Minimum: 10, Maximum: 14}

		actual := MiniMaxSum(input)
		Expect(actual).To(Equal(expected))
	})

	It("1 1 3 4 5 => 9 13", func() {
		input := []int64{1, 1, 3, 4, 5}
		expected := solutions.MiniMax{Minimum: 9, Maximum: 13}

		actual := MiniMaxSum(input)
		Expect(actual).To(Equal(expected))
	})

	It("1 1 3 4 4 => 9 13", func() {
		input := []int64{1, 1, 3, 4, 4}
		expected := solutions.MiniMax{Minimum: 9, Maximum: 12}

		actual := MiniMaxSum(input)
		Expect(actual).To(Equal(expected))
	})

	It("0 0 0 0 0 => 0 0", func() {
		input := []int64{0, 0, 0, 0, 0}
		expected := solutions.MiniMax{Minimum: 0, Maximum: 0}

		actual := MiniMaxSum(input)
		Expect(actual).To(Equal(expected))
	})

	It("5 4 3 2 1 => 10 14", func() {
		input := []int64{5, 4, 3, 2, 1}
		expected := solutions.MiniMax{Minimum: 10, Maximum: 14}

		actual := MiniMaxSum(input)
		Expect(actual).To(Equal(expected))
	})
})
