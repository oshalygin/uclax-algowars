package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("CompareTheTriplets", func() {
	It("Input: [5 6 7] [3 6 10] => [1 1]", func() {
		aliceInput := []int{5, 6, 7}
		bobInput := []int{3, 6, 10}
		expected := []int{1, 1}

		actual := CompareTheTriplets(aliceInput, bobInput)
		Expect(actual).To(Equal(expected))
	})

	It("Input: [1 1 1] [1 1 1] => [0 0]", func() {
		aliceInput := []int{1, 1, 1}
		bobInput := []int{1, 1, 1}
		expected := []int{0, 0}

		actual := CompareTheTriplets(aliceInput, bobInput)
		Expect(actual).To(Equal(expected))
	})

	It("Input: [2 2 2] [1 1 1] => [3 0]", func() {
		aliceInput := []int{2, 2, 2}
		bobInput := []int{1, 1, 1}
		expected := []int{3, 0}

		actual := CompareTheTriplets(aliceInput, bobInput)
		Expect(actual).To(Equal(expected))
	})
})
