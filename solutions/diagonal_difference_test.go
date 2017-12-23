package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("DiagonalDifference", func() {
	It("Input: [[11, 2, 4], [4, 5, 6], [10, 8, -12]] => 15", func() {

		input := [][]int{
			[]int{11, 2, 4}, []int{4, 5, 6}, []int{10, 8, -12},
		}

		expected := 15

		actual := DiagonalDifference(input)
		Expect(actual).To(Equal(expected))
	})

	It("Input: [[11, 2, 4, 3], [4, 5, 6, 3], [10, 8, -12, 3], [1, 1, 1, 1]] => 13", func() {

		input := [][]int{
			[]int{11, 2, 4, 3}, []int{4, 5, 6, 3}, []int{10, 8, -12, 3}, []int{1, 1, 1, 1},
		}

		expected := 13

		actual := DiagonalDifference(input)
		Expect(actual).To(Equal(expected))
	})
})
